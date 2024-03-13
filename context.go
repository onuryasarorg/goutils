package goutils

import (
	"github.com/google/uuid"
)

type Context struct {
	UserId     uuid.UUID
	RoleId     int
	AppId      int
	MerchantId uuid.UUID
	HasId      int
	ProjectId  int
	CustomData string
	Key        string
}
