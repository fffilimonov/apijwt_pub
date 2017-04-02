package models

import (
    "crypto/sha1"
    "encoding/base64"
)

type User struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (user *User) GetPassword() string {
    hasher := sha1.New()
    hasher.Write([]byte(user.Password))
    sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
    return sha
}
