package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string                 `json:"name"`
	Age   int                    `json:"age"`
	Hobby []string               `json:"hobby"`
	Pet   map[string]interface{} `json:"pet"`
}

func (p Person) String() string {
	//return fmt.Sprintf("Person's attribute\nname:%v\nage:%v\nhobby:%v\npet:%v", p.Name, p.Age, p.Hobby, p.Pet)
	return fmt.Sprintf("Person's attribute\nname:%v\nage:%v\nhobby:%v\npet's name:%v\npet's age:%v", p.Name, p.Age, p.Hobby, p.Pet["petName"], p.Pet["petAge"])
}

func main() {
	jsonStr :=
		`{
		"name": "小张",
		"age": 30,
		"hobby": ["唱","跳","rap","篮球"],
		"pet":{"petName":"小花","petAge":3}
	}`
	var p Person
	_ = json.Unmarshal([]byte(jsonStr), &p)
	fmt.Println(p)
}
