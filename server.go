package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println("Hello, World!")

	password := "s3kr1t!"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("hash=%v\n", string(hash))
}
