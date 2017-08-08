package main

import "fmt"

func main() {
	key := genKey()
	for k, v := range key {
		fmt.Printf("\"%s\":\"%s\"\n", string(k), string(v))
	}
}
