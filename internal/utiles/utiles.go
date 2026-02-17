package utiles

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Tokens  interface{} `json:"tokens,omitempty"`
}
