package main

import "fmt"

func main() {
	r, err := GenerateRandomString(52)
	fmt.Println(err)
	fmt.Println(r)
}
