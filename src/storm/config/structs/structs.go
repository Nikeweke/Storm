package structs 

type ResponseMessage struct {
	Success int    `json:"success"`
	Message string `json:"message"`
}

type StringArray map[string]interface{}