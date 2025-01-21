package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

const (
	// DefaultTimeoutSeconds is the number of seconds before a timeout occurs
	// when sending a webhook
	DefaultTimeoutSeconds = 10
)

var client = &http.Client{
	Timeout: time.Second * DefaultTimeoutSeconds,
}

func GetURL(address string, useTLS bool) string {
	if useTLS {
		return "https://" + address
	} else {
		return "http://" + address
	}
}

func SendPostJSON(url string, request, response any) error {
	jsonBody, _ := json.Marshal(request)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = getJsonFromBody(resp.Body, response)
	if err != nil {
		return err
	}

	return nil
}

func SendRegisterRequest(url, hostname, password string) (string, error) {
	resp := &RegisterResponse{}
	err := SendPostJSON(url,
		&RegisterRequest{
			Hostname: hostname,
			Password: password,
		},
		resp,
	)
	if err != nil {
		return "", err
	}
	if resp.Type == "SUCCESS" {
		return resp.Secret, nil
	}
	return "", errors.New(resp.Type)
}

func SendPingRequest(url string) error {
	pingRequest, _ := json.Marshal(&PingRequest{Type: "PING"})
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(pingRequest))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body := new(TypeResponse)
	err = getJsonFromBody(resp.Body, body)
	if err != nil {
		return err
	}
	if body.Type != "PONG" {
		return errors.New("the ping response is not PONG")
	}

	return nil
}

func SendHealthCheckRequest(url, hostname, secret string) error {
	resp := &HealthCheckResponse{}
	err := SendPostJSON(url,
		&HealthCheckRequest{
			Type:     "HEALTH_CHECK",
			Hostname: hostname,
			Secret:   secret,
		},
		resp,
	)
	if err != nil {
		return err
	}

	if resp.Type == "SUCCESS" && resp.Authenticated {
		return nil
	}
	return errors.New("the health check response is not SUCCESS")
}

func SendScreenshot(url, filename, secret string) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return err
	}
	fh, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fh.Close()
	_, err = io.Copy(fw, fh)
	if err != nil {
		return err
	}
	err = writer.Close()
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(body.Bytes()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-Secret", secret)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	return nil
}
