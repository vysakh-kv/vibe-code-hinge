package supabase

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

// SupabaseClient represents a client for interacting with Supabase authentication
type SupabaseClient struct {
	URL    string
	APIKey string
	client *http.Client
}

// SignUpResponse represents the response from a Supabase sign up request
type SignUpResponse struct {
	AccessToken  string        `json:"access_token,omitempty"`
	RefreshToken string        `json:"refresh_token,omitempty"`
	ID           string        `json:"id,omitempty"`
	Aud          string        `json:"aud,omitempty"`
	Role         string        `json:"role,omitempty"`
	Email        string        `json:"email,omitempty"`
	Phone        string        `json:"phone,omitempty"`
	ConfirmationSentAt string  `json:"confirmation_sent_at,omitempty"`
	AppMetadata  map[string]interface{} `json:"app_metadata,omitempty"`
	UserMetadata map[string]interface{} `json:"user_metadata,omitempty"`
	Identities   []map[string]interface{} `json:"identities,omitempty"`
	CreatedAt    string        `json:"created_at,omitempty"`
	UpdatedAt    string        `json:"updated_at,omitempty"`
	IsAnonymous  bool          `json:"is_anonymous,omitempty"`
}

// SignInResponse represents the response from a Supabase sign in request
type SignInResponse struct {
	AccessToken  string        `json:"access_token,omitempty"`
	RefreshToken string        `json:"refresh_token,omitempty"`
	User         *SupabaseUser `json:"user,omitempty"`
	Error        *Error        `json:"error,omitempty"`
}

// SupabaseUser represents a user in Supabase
type SupabaseUser struct {
	ID           string                 `json:"id"`
	Email        string                 `json:"email"`
	UserMetadata map[string]interface{} `json:"user_metadata,omitempty"`
	CreatedAt    string                 `json:"created_at"`
	UpdatedAt    string                 `json:"updated_at"`
}

// Error represents an error returned by the Supabase API
type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status,omitempty"`
}

// NewSupabaseClient creates a new Supabase client
func NewSupabaseClient() (*SupabaseClient, error) {
	url := os.Getenv("SUPABASE_URL")
	if url == "" {
		return nil, errors.New("SUPABASE_URL is not set")
	}

	apiKey := os.Getenv("SUPABASE_KEY")
	if apiKey == "" {
		return nil, errors.New("SUPABASE_ANON_KEY is not set")
	}

	return &SupabaseClient{
		URL:    url,
		APIKey: apiKey,
		client: &http.Client{},
	}, nil
}

// SignUp creates a new user account with email and password
func (c *SupabaseClient) SignUp(email, password string) (*SignUpResponse, error) {
	url := fmt.Sprintf("%s/auth/v1/signup", c.URL)

	reqBody, err := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("apikey", c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("Response body:", string(body))

	var result SignUpResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		errMsg := "Supabase sign up failed"
		if result.Error != nil {
			errMsg = result.Error.Message
		}
		return &result, errors.New(errMsg)
	}

	return &result, nil
}

// SignIn authenticates a user with email and password
func (c *SupabaseClient) SignIn(email, password string) (*SignInResponse, error) {
	url := fmt.Sprintf("%s/auth/v1/token?grant_type=password", c.URL)

	reqBody, err := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("apikey", c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	var result SignInResponse
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return nil, err
	}

	if resp.StatusCode >= 400 {
		errMsg := "Supabase sign in failed"
		if result.Error != nil {
			errMsg = result.Error.Message
		}
		return &result, errors.New(errMsg)
	}

	return &result, nil
}

// VerifyToken verifies a JWT token
func (c *SupabaseClient) VerifyToken(token string) (bool, *SupabaseUser, error) {
	url := fmt.Sprintf("%s/auth/v1/user", c.URL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, nil, err
	}

	req.Header.Set("apikey", c.APIKey)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := c.client.Do(req)
	if err != nil {
		return false, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return false, nil, errors.New("invalid token")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, nil, err
	}

	var user SupabaseUser
	if err := json.Unmarshal(body, &user); err != nil {
		return false, nil, err
	}

	return true, &user, nil
}
