package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	Id    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	p1 := product{Id: 1, Name: "Book", Price: 10.85, Tags: []string{"books", "library"}}
	p1Json, _ := json.Marshal(p1)

	fmt.Println(string(p1Json))

	var p2 product
	jsonString := `{"id":2,"name":"Pen","price":1.15,"tags":["pen","library"]}`
	json.Unmarshal([]byte(jsonString), &p2)

	fmt.Println(p2)
}
