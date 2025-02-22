package http

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

const (
	// DefaultTimeoutSeconds is the number of seconds before a timeout occurs
	DefaultTimeoutSeconds = 30
)

var client = &http.Client{
	Timeout: time.Second * DefaultTimeoutSeconds,
}

func SendScreenshot(filename, url string) error {
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
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	return nil
}
