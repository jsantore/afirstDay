package main

import "fmt"

func main() {
	//names := make([]string, 0,6)
	//for i:= 0; i<8; i++{
	//	append(names, fmt.Sprintf("%d", i))
	//
	//}
	wages := map[string]float32{
		"Ed":  450.17,
		"Ann": 375.99,
	}
	fmt.Print(wages)
}
