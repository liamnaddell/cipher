package main

import "fmt"

func main() {
	key := genKey()
	writeRapToFile("secret.jeff", key)
	final := decodeRapFromFile("secret.jeff")
	printRpay(final)
}

//print map of Rune Array
func printRpay(dag map[rune]rune) {
	for k, v := range dag {
		fmt.Printf("[\"%s\":\"%s\"],", string(k), string(v))
	}
}

//reverse rune array map
func reverserap(rap map[rune]rune) map[rune]rune {
	var revrap = make(map[rune]rune)
	for k, v := range rap {
		revrap[v] = k
	}
	return revrap
}
