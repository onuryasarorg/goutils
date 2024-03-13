package goutils

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	key, secret, bucket, region string
	// ErrNoFilename will return an error if no filename is provided by the user.
	ErrNoFilename = errors.New("no filename provided")
	// ErrNoFilename will return an error if no request type is provided by the user.
	ErrNoRequest = errors.New("no request type provided")
	// ErrNoDuration will return an error if no duration is provided by the user.
	ErrNoDuration = errors.New("no duration provided")
	// ErrNegativeDuration will return an error if a negative duration is provided by the user.
	ErrNegativeDuration = errors.New("negative duration provided")
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(data map[string]interface{}) (*Response, error) {
	jsonData, _ := json.Marshal(data)
	statusCode := http.StatusOK
	if data["status"] != nil && !data["status"].(bool) {
		statusCode = http.StatusNotFound
	}
	return &Response{
		StatusCode: statusCode,
		Headers:    map[string]string{},
		Body:       string(jsonData),
	}, nil
}

const (
	None = iota
	Root
	MerchantAdmin
	Admin
	Superuser
	User
	Member
)
