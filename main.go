package main

import "fmt"

const message = "hello world"

func main() {
	key := genKey()
	msp := []rune(message)
	var nmsg []rune
	for i := 0; i < len(msp); i++ {
		nmsg = append(nmsg, key[msp[i]])
	}
	fmt.Println("encoded:", string(nmsg))
	var dmsg []rune
	revkey := reverserap(key)
	for i := 0; i < len(nmsg); i++ {
		dmsg = append(dmsg, revkey[nmsg[i]])
	}
	fmt.Println("decoded:", string(dmsg))

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
