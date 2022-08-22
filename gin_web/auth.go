package main

import (
	"errors"
)

var UserData map[string]string

func init() {
	UserData = map[string]string{
		"test": "test",
	}
}

// 斷使用者是否存在的程式
func CheckUserIsExist(username string) bool {
	_, isExist := UserData[username]
	return isExist
}

// 密碼比對
func CheckPassword(p1 string, p2 string) error {
	if p1 == p2 {
		return nil
	} else {
		return errors.New("password is no correct")
	}
}

// 組合上面個驗證身份
func Auth(username string, password string) error {
	if isExist := CheckUserIsExist(username); isExist {
		return CheckPassword(UserData[username], password)
	} else {
		return errors.New("user is not exist")
	}
}
