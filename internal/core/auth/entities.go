package auth

import (
	"errors"
	"time"
)

type User struct {
	Login     string
	Email     string
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

var ErrPostNotFound = errors.New("user not found")
