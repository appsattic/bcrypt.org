package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Hello, World!")

	password := "s3kr1t!"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	checkErr(err)

	cost, err := bcrypt.Cost(hash)
	checkErr(err)

	fmt.Printf("Hash=%v\n", string(hash))
	fmt.Printf("Cost=%v\n", cost)
}
