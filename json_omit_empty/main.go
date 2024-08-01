package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	alice := map[string]interface{}{
		"name": "Alice",
		"age":  18,
	}
	bs, err := json.Marshal(Person{
		Name: alice["name"].(string),
		Age:  alice["age"].(int),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("serialized person without omitting empty fields: \t%s\n", bs)

	bs, err = json.Marshal(PersonOmitEmpty{
		Name: alice["name"].(string),
		Age:  alice["age"].(int),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("serialized person omitting empty fields: \t\t%s\n", bs)
}

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Alias string `json:"alias"`
}

type PersonOmitEmpty struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Alias string `json:"alias,omitempty"`
}
