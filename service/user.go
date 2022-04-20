package service

import (
	"Cloudflare-Assistant/model"
	"sync"
)

var (
	User     *model.User
	UserLock sync.RWMutex
)

func CheckValid() (bool, error) {
	return true, nil
}
