package customerror

import "fmt"

type ValidationError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
}

type CustomError struct {
	Code            int               `json:"code"`
	Message         string            `json:"message"`
	ValidationError []ValidationError `json:"validationError,omitempty"`
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("Code: %d Message: %s", c.Code, c.Message)
}
