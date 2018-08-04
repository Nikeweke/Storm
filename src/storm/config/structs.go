package config 

type ResponseMessage struct {
	Success int    `json:"success"`
	Message string `json:"message"`
}