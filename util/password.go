package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Get256Pw(string2 string) string {
	PasswordRecived := string2
	last3 := PasswordRecived[len(PasswordRecived)-3:]
	h := sha256.New()
	h.Write([]byte(PasswordRecived + last3))
	hashString := hex.EncodeToString(h.Sum(nil))
	//hashString := "123456"
	fmt.Println(hashString)
	return hashString
}
