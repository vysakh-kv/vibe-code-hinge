package utils

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Response represents a JSON response structure
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SuccessResponse creates a success response
func SuccessResponse(data interface{}, message string) Response {
	return Response{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}

// ErrorResponse creates an error response
func ErrorResponse(message string) Response {
	return Response{
		Status: "error",
		Error:  message,
	}
}

// RespondWithJSON writes a JSON response
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// RespondWithError writes an error response
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, ErrorResponse(message))
}

// GetPathParam gets a path parameter from the URL
func GetPathParam(r *http.Request, param string) string {
	return mux.Vars(r)[param]
}

// GetPathParamInt gets an integer path parameter from the URL
func GetPathParamInt(r *http.Request, param string) (int64, error) {
	return strconv.ParseInt(GetPathParam(r, param), 10, 64)
}

// GetQueryParam gets a query parameter from the URL
func GetQueryParam(r *http.Request, param string) string {
	return r.URL.Query().Get(param)
}

// GetQueryParamInt gets an integer query parameter from the URL
func GetQueryParamInt(r *http.Request, param string, defaultValue int) int {
	valueStr := GetQueryParam(r, param)
	if valueStr == "" {
		return defaultValue
	}
	
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	
	return value
}

// DecodeJSONBody decodes a JSON request body into a destination struct
func DecodeJSONBody(r *http.Request, dst interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(dst)
} 