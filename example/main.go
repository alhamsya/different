package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/alhamsya/different"
)

type User struct {
	Name    string
	Age     int `diff:"-"`
	Address Address
	Telephone
}

type Address struct {
	City   string
	Street string
}

type Telephone struct {
	Type        string
	PhoneNumber string `diff:"phone_number"`
}

func main() {
	origin := &User{
		Name: "Alhamsya",
		Age:  23,
		Address: Address{
			City:   "Kediri",
			Street: "New street",
		},
		Telephone: Telephone{
			Type:        "Rumah",
			PhoneNumber: "081",
		},
	}

	newData := &User{
		Name: "Bintang",
		Age:  25,
		Address: Address{
			City:   "Kediri",
			Street: "Old Street",
		},
		Telephone: Telephone{
			Type:        "Rumah",
			PhoneNumber: "085",
		},
	}

	diff, err := different.GenerateDiff(origin, newData)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(diff))
}

func printData(result interface{}) {
	src, _ := json.Marshal(result)

	dst := &bytes.Buffer{}
	if err := json.Indent(dst, src, "", "  "); err != nil {
		panic(err)
	}

	fmt.Println(dst.String())
}
