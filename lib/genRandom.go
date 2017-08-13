package cipher

import (
	"crypto/rand"
	"math/big"
)

func genRandomString(s int) (string, error) {
	var rs string
	var previous []string
	var c bool
	for i := 0; i < s+1; i++ {
		st, err := getToken(1)
		if err != nil {
			return "", err
		}
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
	return rs, nil
}

// thank you stackOverflow question asker
func getToken(length int) (string, error) {
	token := ""
	//codeAlphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//codeAlphabet += "abcdefghijklmnopqrstuvwxyz"
	//codeAlphabet += "0123456789"
	codeAlphabet := clear

	for i := 0; i < length; i++ {
		n, err := cryptoRandSecure(int64(len(codeAlphabet)))
		if err != nil {
			return "", err
		}
		token += string(codeAlphabet[n])
	}
	return token, nil
}

func cryptoRandSecure(max int64) (int64, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return 0, err
	}
	return nBig.Int64(), nil
}
