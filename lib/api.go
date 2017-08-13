package cipher

import "fmt"
import "io/ioutil"

//wraps up the api for the cli to use

func NewKey(keyfile string) error {
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

func NewFileEncode(messagefile string, keyfile string) error {
	message, err := ioutil.ReadFile(messagefile)
	if err != nil {
		return err
	}
	//printRap(key)
	msp := string(message)
	return NewEncode(msp, keyfile)
}
func NewEncode(message string, keyfile string) error {
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

func NewDecode(message string, keyfile string) error {
	var err error
	key, err := decodeRapFromFile(keyfile)
	if err != nil {
		return err
	}
	//printRap(key)
	var dmsg []rune
	var msp = []rune(message)
	revkey := Reverserap(key)
	for i := 0; i < len(message); i++ {
		dmsg = append(dmsg, revkey[msp[i]])
	}
	fmt.Println("decoded:", string(dmsg))
	return err
}

func NewFileDecode(messagefile string, keyfile string) error {
	nmessage, err := ioutil.ReadFile(messagefile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var rmessage = string(nmessage)
	return NewDecode(rmessage, keyfile)
}

func PrintRap(dag map[rune]rune) {
	for k, v := range dag {
		fmt.Printf("[\"%s\":\"%s\"],", string(k), string(v))
	}
}

//reverse rune array map
func Reverserap(rap map[rune]rune) map[rune]rune {
	var revrap = make(map[rune]rune)
	for k, v := range rap {
		revrap[v] = k
	}
	return revrap
}
