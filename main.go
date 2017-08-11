package main

import "fmt"

func main() {
	fmt.Println("hello world")

}

//print map of Rune Array
func printRpay(dag map[rune]rune) {
	for k, v := range dag {
		fmt.Printf("\"%s\":\"%s\"\n", string(k), string(v))
	}
}
