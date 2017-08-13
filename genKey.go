package main

import (
	"fmt"
	"strings"
)

const clear = "1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM,./`!@#$%^&*()-=~_+<> ?;':\"\\"

var code string

var clearArray []rune
var codeArray []rune

func init() {
	clearSplit := strings.Split(clear, "")
	clearArray = strayToRay(clearSplit)
	code, err := genKeyCode()
	if err != nil {
		panic(err)
	}
	codeSplit := strings.Split(code, "")
	codeArray = strayToRay(codeSplit)
}

//string array to rune array
func strayToRay(split []string) (array []rune) {
	for i := 0; i < len(split); i++ {
		array = append(array, []rune(split[i])...)
	}
	return
}

func genKey() map[rune]rune {
	fmt.Println(len(clear))
	fmt.Println(len(code))
	var i int
	var key = make(map[rune]rune)
	for {
		if i <= len(clear) {
			for o := 0; o < len(clearArray); o++ {
				char := clearArray[o]
				key[char] = codeArray[o]
			}
		} else {
			break
		}
		i++
	}
	return key
}
