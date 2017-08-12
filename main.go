package main

import "fmt"

func main() {
	/*key := genKey()
	//writeRapToFile("secret.jeff", key)
	//final := decodeRapFromFile("secret.jeff")
	//printRpay(final)
	msp := []rune("hello world")
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
	*/
	startCli()
}

//print map of rune
func printRap(dag map[rune]rune) {
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
