package apiUtil

type Field struct {
	Status  int
	Message string
	Data    interface{}
	Errors  interface{}
}

type JsonResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}
