package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Francis Seneve")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer res.Body.Close()

	fmt.Println("Http response status: ", res.Status)
}
