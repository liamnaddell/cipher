package main

import "fmt"

func main() {
	fmt.Println("hello world")
	printRpay(genKey())
}

//print map of Rune Array
func printRpay(dag map[rune]rune) {
	for k, v := range dag {
		fmt.Printf("[\"%s\":\"%s\"],", string(k), string(v))
	}
}
