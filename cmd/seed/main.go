package main

import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
)

func main() {
	hash, err := bcrypt.GenerateFromPassword([]byte("this is my admin password"), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(hash))
}
