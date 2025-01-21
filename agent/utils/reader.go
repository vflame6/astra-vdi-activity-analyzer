package utils

import (
	"golang.org/x/term"
	"strings"
	"syscall"
)

func ReadPassword() (string, error) {
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	password := string(bytePassword)

	return strings.TrimSpace(password), nil
}
