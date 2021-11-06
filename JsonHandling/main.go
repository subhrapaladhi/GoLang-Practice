package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	mycourses := []course{
		{"golang", 100, "youtube", "123", []string{"go", "code"}},
		{"c++", 200, "udemy", "123", []string{"c++", "code"}},
		{"java", 300, "coursera", "123", []string{"java", "code"}},
	}

	finalJson, err := json.MarshalIndent(mycourses, "", " ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonData := []byte(`
	{
		"name": "c++",
		"price": 200,
		"platform": "udemy",
		"tags": ["c++", "code"]
	 }
	`)

	var mycourse course

	if json.Valid(jsonData) {
		json.Unmarshal(jsonData, &mycourse)
		fmt.Printf("%#v\n", mycourse)
	} else {
		fmt.Println("invalid json")
	}

	var myData map[string]interface{}
	json.Unmarshal(jsonData, &myData)

	for k, v := range myData {
		fmt.Println(k, v)
	}
}
