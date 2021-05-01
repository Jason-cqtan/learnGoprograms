package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type person struct {
	First string
}

func main() {
	p1 := person{
		"jason",
	}
	bs, err := toJson(p1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println((string)(bs))
}

func toJson(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		//return []byte{}, fmt.Errorf("There was an error in toJSON: %v", err)
		//return []byte{}, fmt.Errorf("There was an error marshalling %v in toJSON: %v", a, err)
		return []byte{}, errors.New(fmt.Sprintf("There was an error in toJSON %v", err))
	}

	return bs, err
}
