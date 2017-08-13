package main

import "fmt"
import "io/ioutil"

//wraps up the api for the cli to use

func newKey(keyfile string) error {
	var err error
	key := genKey()
	// here
	writeRapToFile(keyfile, key)
	return err
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

func newFileEncode(messagefile string, keyfile string) error {
	message, err := ioutil.ReadFile(messagefile)
	if err != nil {
		return err
	}
	//printRap(key)
	msp := string(message)
	return newEncode(msp, keyfile)
}
func newEncode(message string, keyfile string) error {
	key, err := decodeRapFromFile(keyfile)
	if err != nil {
		return err
	}
	msp := []rune(message)
	var nmsg []rune
	for i := 0; i < len(msp); i++ {
		nmsg = append(nmsg, key[msp[i]])
	}
	fmt.Println("encoded:", string(nmsg))
	return err
}

func newDecode(message string, keyfile string) error {
	var err error
	key, err := decodeRapFromFile(keyfile)
	if err != nil {
		return err
	}
	//printRap(key)
	var dmsg []rune
	var msp = []rune(message)
	revkey := reverserap(key)
	for i := 0; i < len(message); i++ {
		dmsg = append(dmsg, revkey[msp[i]])
	}
	fmt.Println("decoded:", string(dmsg))
	return err
}

func newFileDecode(messagefile string, keyfile string) error {
	nmessage, err := ioutil.ReadFile(messagefile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var rmessage = string(nmessage)
	return newDecode(rmessage, keyfile)
}
