package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	bytearray, err := ioutil.ReadFile("sillyRec.txt")
	if err != nil {
		log.Fatal("Help file error:", err)
	}
	str := string(bytearray)
	lines := strings.Split(str, "\n")
	for position, line := range lines {
		if position%2 == 0 {
			fmt.Println(line)
		}
	}
}
