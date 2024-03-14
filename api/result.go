package api

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

func ErrMessage(status bool, err error) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": err}
}

func RespondError(err error) (*Response, error) {
	return &Response{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{},
		Body:       "",
	}, err
}

func Respond(data map[string]interface{}) (*Response, error) {
	jsonData, _ := json.Marshal(data)
	return &Response{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{},
		Body:       string(jsonData),
	}, nil
}

func RespondNoContent() (*Response, error) {
	return &Response{
		StatusCode: http.StatusNoContent,
		Headers:    map[string]string{},
		Body:       "",
	}, nil
}

func ResMessage(status bool, message string) (bool, Response) {
	jsonData, _ := json.Marshal(Message(status, message))
	return status, Response{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{},
		Body:       string(jsonData),
	}
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
