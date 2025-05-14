package models

// SuccessResponse represents a successful API response
type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse represents an error API response
type ErrorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

// NewSuccessResponse creates a new success response
func NewSuccessResponse(message string, data interface{}) SuccessResponse {
	return SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}

// NewErrorResponse creates a new error response
func NewErrorResponse(error string) ErrorResponse {
	return ErrorResponse{
		Status: "error",
		Error:  error,
	}
}

// Pagination represents pagination metadata
type Pagination struct {
	Total       int `json:"total"`
	PageSize    int `json:"page_size"`
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
}

// PaginatedResponse represents a paginated API response
type PaginatedResponse struct {
	Status     string      `json:"status"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

// NewPaginatedResponse creates a new paginated response
func NewPaginatedResponse(data interface{}, pagination Pagination) PaginatedResponse {
	return PaginatedResponse{
		Status:     "success",
		Data:       data,
		Pagination: pagination,
	}
}
