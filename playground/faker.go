package main

import (
	"github.com/google/uuid"
)

func newUuid() string {
	return uuid.New().String()
}
