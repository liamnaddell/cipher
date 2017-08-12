package main

import (
	"crypto/rand"
	"log"
	"math/big"
)

func genRandomString(s int) string {
	var rs string
	var previous []string
	var c bool
	for i := 0; i < s+1; i++ {
		st := getToken(1)
		for q := 0; q < len(previous); q++ {
			if st == previous[q] {
				c = true
			}
		}
		if c {
			i -= 1
			c = false
			continue
		}
		rs += st
		previous = append(previous, st)
	}
	return rs
}

// thank you stackOverflow question asker
func getToken(length int) string {
	token := ""
	//codeAlphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//codeAlphabet += "abcdefghijklmnopqrstuvwxyz"
	//codeAlphabet += "0123456789"
	codeAlphabet := clear

	for i := 0; i < length; i++ {
		token += string(codeAlphabet[cryptoRandSecure(int64(len(codeAlphabet)))])
	}
	return token
}

func cryptoRandSecure(max int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		log.Println(err)
	}
	return nBig.Int64()
}
