package main

import "encoding/gob"
import "bytes"
import "io/ioutil"
import "log"

func encodeMap(rap map[rune]rune) *bytes.Buffer {
	var bay = new(bytes.Buffer)
	enc := gob.NewEncoder(bay)
	_ = enc.Encode(rap)
	return bay
}

func writeRapToFile(filename string, rap map[rune]rune) {
	emap := encodeMap(rap)
	ioutil.WriteFile(filename, emap.Bytes(), 0777)
}

func decodeMap(buf *bytes.Buffer) map[rune]rune {
	var dm = make(map[rune]rune)
	decoder := gob.NewDecoder(buf)
	_ = decoder.Decode(&dm)
	return dm
}

func decodeRapFromFile(filename string) map[rune]rune {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	buf := bytes.NewBuffer(file)
	return decodeMap(buf)
}
