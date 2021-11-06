package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	res, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println(res.StatusCode)
	fmt.Println(res.ContentLength)

	var resString strings.Builder
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%T\n", content)
	// fmt.Println(content)
	// fmt.Println(string(content))

	byteCount, err := resString.Write(content)
	if err != nil {
		panic(err)
	}
	fmt.Println(byteCount)
	fmt.Println(resString.String())
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	reqBody := strings.NewReader(`
	{
		"name": "go lang",
		"platform": "youtube"
	}
	`)

	res, err := http.Post(myurl, "application/json", reqBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("name", "subhra")
	data.Add("age", "21")

	res, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
