package cipher

import "encoding/gob"
import "bytes"
import "io/ioutil"

func encodeMap(rap map[rune]rune) (*bytes.Buffer, error) {
	var bay = new(bytes.Buffer)
	enc := gob.NewEncoder(bay)
	err := enc.Encode(rap)
	return bay, err
}

func writeRapToFile(filename string, rap map[rune]rune) error {
	emap, err := encodeMap(rap)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, emap.Bytes(), 0777)
	return err
}

func decodeMap(buf *bytes.Buffer) (map[rune]rune, error) {
	var dm = make(map[rune]rune)
	decoder := gob.NewDecoder(buf)
	err := decoder.Decode(&dm)
	return dm, err
}

func decodeRapFromFile(filename string) (map[rune]rune, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(file)
	return decodeMap(buf)
}
