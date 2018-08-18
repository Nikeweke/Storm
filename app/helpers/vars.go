package helpers 

type ResponseMessage struct {
	Success int    `json:"success"`
	Message string `json:"message"`
}

type StringArray map[string]interface{}

type Route struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	Name   string `json:"name"`
} 