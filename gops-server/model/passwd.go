package model

import (
	"golang.org/x/crypto/scrypt"
	"bytes"
)

const (
	N        = 16384
	R        = 8
	P        = 1
	KEYLENTH = 32
)

func EncryptPassword(user, pwd, ctime string) ([]byte, error) {
	c, err := scrypt.Key([]byte(pwd), []byte(ctime+user), N, R, P, KEYLENTH)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func CheckPassword(user, pwd, ctime, cpwd string) bool {
	c, err := EncryptPassword(user, pwd, ctime)
	if err != nil || !bytes.Equal(c, []byte(cpwd)) {
		return false
	}
	return true
}
