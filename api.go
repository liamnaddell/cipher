package main

import "fmt"
import "io/ioutil"

//wraps up the api for the cli to use

func newKey(keyfile string) error {
	var err error
	key := genKey()
	writeRapToFile(keyfile, key)
	return err
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

func newEncode(message string, keyfile string) error {
	var err error
	key := decodeRapFromFile(keyfile)
	//printRap(key)
	msp := []rune(message)
	var nmsg []rune
	for i := 0; i < len(msp); i++ {
		nmsg = append(nmsg, key[msp[i]])
	}
	fmt.Println("encoded:", string(nmsg))
	return err
}

func newDecode(messagefile string, keyfile string) error {
	var err error
	nmessage, err := ioutil.ReadFile(messagefile)
	if err != nil {
		return err
	}
	var message = []rune(string(nmessage))
	key := decodeRapFromFile(keyfile)
	//printRap(key)
	var dmsg []rune
	revkey := reverserap(key)
	for i := 0; i < len(message); i++ {
		dmsg = append(dmsg, revkey[message[i]])
	}
	fmt.Println("decoded:", string(dmsg))
	return err
}
