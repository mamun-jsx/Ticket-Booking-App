package httpresponse

// import "strings"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// func NewError(code int, message string, details ...string) *Error {
// 	detailsStr := ""
// 	if len(details) > 0 && details[0] != "" {
// 		detailsStr = strings.Join(details, " ")
// 	}
// 	return &Error{
// 		Code:    code,
// 		Message: message,
// 		Details: detailsStr,
// 	}
// }
