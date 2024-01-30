package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("https://news.ycombinator.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	fmt.Print(response.Body)
}
