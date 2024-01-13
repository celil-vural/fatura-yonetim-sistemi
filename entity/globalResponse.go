package entity

type GlobalResponse struct {
	Data        interface{}   `json:"data"`
	ErrorDetail ErrorResponse `json:"error_details"`
	Success     bool          `json:"success"`
}
type ErrorResponse struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}
