package common

import (
	"github.com/satori/go.uuid"
	"strings"
)

func NewUuidV1() string {
	u1 := uuid.NewV1()
	s := u1.String()
	return strings.Replace(s,"-","",-1)
}