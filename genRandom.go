package main

import (
	"crypto/rand"
	//"encoding/base64"
	"fmt"
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
				fmt.Println("continue, ", st)
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

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
/*
func genRandomString(s int) string {
	var rs string
	var lnb []string
	for i := 0; i < s+1; i++ {
		b, err := genRandomBytes(1)
		if err != nil {
			log.Fatal(err)
		}
		nb := base64.StdEncoding.EncodeToString(b)
		tnb, _ := base64.StdEncoding.DecodeString(nb)
		lnb = append(lnb, nb)
		for q := 0; q < len(lnb); q++ {
			if nb == lnb[q] {
				fmt.Println("Continnue:", nb)
				continue
			}
		}
		rs = rs + string(tnb)
	}
	fmt.Println(lnb)
	return rs
}
*/
