package router

type PingRequest struct {
	Type string `json:"type"`
}

type HealthRequest struct {
	Type     string `json:"type"`
	Hostname string `json:"hostname"`
	Secret   string `json:"secret"`
}

type RegisterRequest struct {
	Hostname string `json:"hostname"`
	Password string `json:"password"`
}
