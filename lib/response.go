package lib

type Response struct {
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
}
