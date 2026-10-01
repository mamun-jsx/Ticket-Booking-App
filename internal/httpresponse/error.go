package httpresponse

// Error represents the standardized JSON structure for API error responses.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}
