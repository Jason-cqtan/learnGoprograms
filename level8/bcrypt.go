package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := `123`
	bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pass)
	fmt.Println(bs)

	loginPass := `1234`

	err = bcrypt.CompareHashAndPassword(bs,[]byte(loginPass))
	if err != nil {
		fmt.Println("pass not correct")
		return
	}


	fmt.Println("correct pass")

}
