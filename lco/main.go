package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var url = "https://lco.dev"

func main() {

	fmt.Println("NEW LCO")

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	byteData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	content := string(byteData)

	fmt.Println(content)
}
