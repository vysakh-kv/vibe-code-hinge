package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

// PostmanCollection represents the top-level structure of a Postman collection
type PostmanCollection struct {
	Info  Info    `json:"info"`
	Item  []Folder `json:"item"`
}

// Info contains collection metadata
type Info struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Schema      string    `json:"schema"`
	PostmanID   string    `json:"_postman_id"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// Folder represents a folder in the Postman collection
type Folder struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Item        []Item `json:"item"`
}

// Item represents a request in the Postman collection
type Item struct {
	Name     string   `json:"name"`
	Request  Request  `json:"request"`
	Response []string `json:"response"`
}

// Request represents a Postman request
type Request struct {
	Method      string      `json:"method"`
	Header      []Header    `json:"header"`
	URL         URL         `json:"url"`
	Body        *Body       `json:"body,omitempty"`
	Description string      `json:"description,omitempty"`
}

// Header represents an HTTP header
type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

// URL represents a URL in Postman format
type URL struct {
	Raw      string   `json:"raw"`
	Protocol string   `json:"protocol"`
	Host     []string `json:"host"`
	Path     []string `json:"path"`
	Query    []Query  `json:"query,omitempty"`
}

// Query represents a URL query parameter
type Query struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Body represents a request body
type Body struct {
	Mode    string    `json:"mode"`
	Raw     string    `json:"raw,omitempty"`
	Options *Options  `json:"options,omitempty"`
}

// Options represents body options
type Options struct {
	Raw Options2 `json:"raw"`
}

// Options2 represents language options for raw body
type Options2 struct {
	Language string `json:"language"`
}

// parseCurlCommand parses a curl command into its components
func parseCurlCommand(curlCmd string, name string) Item {
	// Default item
	item := Item{
		Name:     name,
		Response: []string{},
	}

	// Extract method
	methodRegex := regexp.MustCompile(`-X\s+([A-Z]+)`)
	method := "GET" // Default method
	methodMatches := methodRegex.FindStringSubmatch(curlCmd)
	if len(methodMatches) > 1 {
		method = methodMatches[1]
	}

	// Extract URL
	urlRegex := regexp.MustCompile(`"(\${BASE_URL}[^"]+)"`)
	urlMatches := urlRegex.FindStringSubmatch(curlCmd)
	urlStr := ""
	if len(urlMatches) > 1 {
		urlStr = urlMatches[1]
		urlStr = strings.Replace(urlStr, "${BASE_URL}", "{{base_url}}", -1)
	}

	// Extract headers
	headerRegex := regexp.MustCompile(`-H\s+"([^"]+)"`)
	headerMatches := headerRegex.FindAllStringSubmatch(curlCmd, -1)
	headers := []Header{}
	for _, match := range headerMatches {
		if len(match) > 1 {
			parts := strings.SplitN(match[1], ":", 2)
			if len(parts) == 2 {
				headers = append(headers, Header{
					Key:   strings.TrimSpace(parts[0]),
					Value: strings.TrimSpace(parts[1]),
					Type:  "text",
				})
			}
		}
	}

	// Extract body
	var body *Body
	bodyRegex := regexp.MustCompile(`-d\s+'([^']+)'`)
	bodyMatches := bodyRegex.FindStringSubmatch(curlCmd)
	if len(bodyMatches) < 2 {
		bodyRegex = regexp.MustCompile(`-d\s+'([^']+)'`)
		bodyMatches = bodyRegex.FindStringSubmatch(curlCmd)
	}
	if len(bodyMatches) < 2 {
		// Try with double quotes
		bodyRegex = regexp.MustCompile(`-d\s+"([^"]+)"`)
		bodyMatches = bodyRegex.FindStringSubmatch(curlCmd)
	}
	if len(bodyMatches) < 2 {
		// Try multiline with backticks (which we convert in the markdown)
		bodyRegex = regexp.MustCompile(`-d\s+'\{([^}]*)\}'`)
		bodyMatches = bodyRegex.FindStringSubmatch(curlCmd)
	}
	if len(bodyMatches) > 1 {
		bodyContent := bodyMatches[1]
		body = &Body{
			Mode: "raw",
			Raw:  bodyContent,
			Options: &Options{
				Raw: Options2{
					Language: "json",
				},
			},
		}
	}

	// Parse URL into components
	urlObj := URL{}
	if urlStr != "" {
		// Replace any path parameters with Postman variables
		urlStr = strings.Replace(urlStr, "{id}", "{{id}}", -1)
		urlStr = strings.Replace(urlStr, "{profile_id}", "{{profile_id}}", -1)
		urlStr = strings.Replace(urlStr, "{match_id}", "{{match_id}}", -1)
		urlStr = strings.Replace(urlStr, "{user_id}", "{{user_id}}", -1)

		urlObj.Raw = urlStr
		urlObj.Protocol = "http"

		// Split URL into host and path
		basePart := "{{base_url}}"
		hostPart := []string{"{{base_url}}"}
		
		// Extract path
		pathPart := strings.TrimPrefix(urlStr, basePart)
		pathPart = strings.TrimPrefix(pathPart, "/")
		pathSegments := []string{}
		if pathPart != "" {
			// Split by / but handle variables properly
			segments := strings.Split(pathPart, "/")
			for _, segment := range segments {
				pathSegments = append(pathSegments, segment)
			}
		}
		urlObj.Host = hostPart
		urlObj.Path = pathSegments

		// Extract query parameters
		if strings.Contains(urlStr, "?") {
			queryPart := strings.Split(urlStr, "?")[1]
			queryParams := strings.Split(queryPart, "&")
			queries := []Query{}
			for _, param := range queryParams {
				if strings.Contains(param, "=") {
					parts := strings.SplitN(param, "=", 2)
					key := parts[0]
					value := ""
					if len(parts) > 1 {
						value = parts[1]
					}
					queries = append(queries, Query{
						Key:   key,
						Value: value,
					})
				}
			}
			urlObj.Query = queries
		}
	}

	// Create request
	request := Request{
		Method: method,
		Header: headers,
		URL:    urlObj,
		Body:   body,
	}

	item.Request = request
	return item
}

func main() {
	// Open the markdown file
	file, err := os.Open("curl_commands.md")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Initialize collection
	collection := PostmanCollection{
		Info: Info{
			Name:        "Vibe Dating App API",
			Description: "API collection for the Vibe Dating App",
			Schema:      "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
			PostmanID:   fmt.Sprintf("vibe-%d", time.Now().Unix()),
			UpdatedAt:   time.Now(),
		},
		Item: []Folder{},
	}

	var currentFolder *Folder
	var currentCurlBlock []string
	var currentRequestName string
	inCurlBlock := false

	// Process the file line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Check if this line starts a new section
		if strings.HasPrefix(line, "## ") {
			// If we were processing a curl block, finalize it
			if inCurlBlock && len(currentCurlBlock) > 0 {
				curlCmd := strings.Join(currentCurlBlock, "\n")
				item := parseCurlCommand(curlCmd, currentRequestName)
				currentFolder.Item = append(currentFolder.Item, item)
				currentCurlBlock = []string{}
				currentRequestName = ""
			}

			// Start a new folder
			folderName := strings.TrimPrefix(line, "## ")
			currentFolder = &Folder{
				Name:        folderName,
				Description: "",
				Item:        []Item{},
			}
			collection.Item = append(collection.Item, *currentFolder)
			inCurlBlock = false
		}

		// Check if this line is a request name (# Comment line inside a code block)
		if inCurlBlock && strings.HasPrefix(line, "# ") {
			currentRequestName = strings.TrimPrefix(line, "# ")
		}

		// Check if this line starts a curl code block
		if strings.HasPrefix(line, "```bash") {
			inCurlBlock = true
			continue
		}

		// Check if this line ends a code block
		if line == "```" && inCurlBlock {
			// Process the curl block we just finished
			if len(currentCurlBlock) > 0 && currentFolder != nil {
				curlCmd := strings.Join(currentCurlBlock, "\n")
				if currentRequestName == "" {
					// Use the first few words as the name
					parts := strings.Split(currentCurlBlock[0], " ")
					if len(parts) > 2 {
						currentRequestName = strings.Join(parts[1:3], " ")
					} else {
						currentRequestName = "Request"
					}
				}
				item := parseCurlCommand(curlCmd, currentRequestName)
				currentFolder.Item = append(currentFolder.Item, item)
				currentCurlBlock = []string{}
				currentRequestName = ""
			}
			inCurlBlock = false
			continue
		}

		// If we're in a curl block, add this line to the current block
		if inCurlBlock && !strings.HasPrefix(line, "```") {
			currentCurlBlock = append(currentCurlBlock, line)
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Write the collection to a file
	outputFile, err := os.Create("vibe_dating_app.postman_collection.json")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	encoder := json.NewEncoder(outputFile)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(collection)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println("Successfully generated Postman collection: vibe_dating_app.postman_collection.json")
} 