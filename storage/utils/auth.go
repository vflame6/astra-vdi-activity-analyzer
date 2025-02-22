package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

const secretKeyLength = 32

func GenerateSecret() string {
	b := make([]byte, secretKeyLength)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}
