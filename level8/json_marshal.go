package main

import (
	"encoding/json"
	"fmt"
)

type simpleUser struct {
	First string
	Age   int
}

func main() {
	u1 := simpleUser{
		First: "James",
		Age:   32,
	}

	u2 := simpleUser{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := simpleUser{
		First: "M",
		Age:   54,
	}

	users := []simpleUser{u1, u2, u3}

	fmt.Println(users)

	jsonByte, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonByte))
}
