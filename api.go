package main

//import "fmt"

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
