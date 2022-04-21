package model

import "time"

type User struct {
	Common
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"is_admin,omitempty"`

	Token          string    `gorm:"UNIQUE_INDEX" json:"-"`
	TokenExpiredAt time.Time `json:"token_expired_at,omitempty"`
}
