package model

import (
	"net/http"
	"time"
)

type (
	User struct {
		CreatedAt   time.Time `json:"created_at"`
		DisplayName string    `json:"display_name"`
		Email       string    `json:"email"`
	}
	UserList map[string]User
)
type CreateUserRequest struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}
type UpdateUserRequest struct {
	DisplayName string `json:"display_name"`
}

func (c *UpdateUserRequest) Bind(r *http.Request) error { return nil }
func (c *CreateUserRequest) Bind(r *http.Request) error { return nil }
