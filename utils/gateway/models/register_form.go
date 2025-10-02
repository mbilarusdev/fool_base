package gateway_models

type RegisterForm struct {
	ServiceName       string    `json:"service_name"`
	InstanceID        string    `json:"instance_id"`
	Endpoints         *Endpoint `json:"resolvers"`
	GatewayAddr       string
	GatewayProtocol   string
	ServiceRemoteAddr string `json:"service_remote_addr"`
	ServiceProtocol   string `json:"service_protocol"`
}
