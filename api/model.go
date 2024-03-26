package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type Response struct {
	// StatusCode is the http code that will be returned back to the user.
	StatusCode int `json:"statusCode,omitempty"`
	// Headers is the information about the type of data being returned back.
	Headers map[string]string `json:"headers,omitempty"`
	// Body will contain the presigned url to upload or download files.
	Body string `json:"body,omitempty"`
}

type Context struct {
	UserId     uuid.UUID `json:"userId"`
	RoleId     int       `json:"roleId"`
	AppId      int       `json:"appId"`
	MerchantId uuid.UUID `json:"merchantId"`
	HasId      bool      `json:"hasId"`
	ProjectId  int       `json:"projectId"`
	CustomData string    `json:"customData"`
}

type Token struct {
	UserId     uuid.UUID
	RoleId     int
	AppId      int
	MerchantId uuid.UUID
	HasId      bool
	ProjectId  int
	CustomData string
	jwt.StandardClaims
}
