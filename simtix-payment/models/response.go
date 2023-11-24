package models

type ResponseBody struct {
	Message string       `json:"message"`
	Data    *interface{} `json:"data,omitempty"`
}

type Response struct {
	Code int
	Body ResponseBody
}
