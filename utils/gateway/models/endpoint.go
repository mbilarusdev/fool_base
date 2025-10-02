package gateway_models

type Endpoint struct {
	Protocol string `json:"protocol"`
	Path     string `json:"path"`
	Method   string `json:"method"`
}
