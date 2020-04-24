package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	p := "admin"
	newPassword, _ := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	fmt.Println(string(newPassword))
}