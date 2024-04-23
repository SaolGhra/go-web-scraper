package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("https://www.ign.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bytes))
}
