package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://loc.dev:3000/learn?coursename=reactjs&paymentid=dsaf3323"

func main() {
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Println(qparams)

	for key, val := range qparams {
		fmt.Println(key, val[0])
		fmt.Printf("%T\n", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "loc.dev",
		Path:    "/tutcss",
		RawPath: "user=subhra",
	}

	url2 := partsOfUrl.String()
	fmt.Println(url2)
}
