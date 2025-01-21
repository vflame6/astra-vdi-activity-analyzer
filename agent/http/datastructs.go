package http

import (
	"encoding/json"
	"io"
)

type PingRequest struct {
	Type string `json:"type"`
}

type TypeResponse struct {
	Type string `json:"type"`
}

type HealthCheckRequest struct {
	Type     string `json:"type"`
	Hostname string `json:"hostname"`
	Secret   string `json:"secret"`
}

type HealthCheckResponse struct {
	Type          string `json:"type"`
	Authenticated bool   `json:"authenticated"`
}

type RegisterRequest struct {
	Hostname string `json:"hostname"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Type   string `json:"type"`
	Secret string `json:"secret"`
}

func getJsonFromBody(body io.ReadCloser, v interface{}) error {
	return json.NewDecoder(body).Decode(v)
}
