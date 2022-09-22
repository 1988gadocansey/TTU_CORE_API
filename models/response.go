package models

type Response struct {
	Code int         `json:"code,omitempty"`
	Body interface{} `json:"body,omitempty"`
	//Data    *fiber.Map  `json:"data"`
	Title   string `json:"title,omitempty"`
	Message string `json:"message,omitempty"`
}
