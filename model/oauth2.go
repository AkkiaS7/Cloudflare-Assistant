package model

type OAuth2 struct {
	Common
	UserId uint64 `json:"user_id"`
	Method string `json:"method"`
}
