package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type CaesarJSON struct {
	Text string `json:"text"`
	Key  int    `json:"key"`
}

func (c *CaesarJSON) Unmarshal(filename string) {
	buff, err := ioutil.ReadFile("caesar/testcases/" + filename)
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}
	err = json.Unmarshal(buff, c)
	if err != nil {
		log.Fatal("Error unmarshaling file: ", err)
	}
}

type VigenereJSON struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}

func (v *VigenereJSON) Unmarshal(filename string) {
	buff, err := ioutil.ReadFile("vigenere/testcases/" + filename)
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}
	err = json.Unmarshal(buff, v)
	if err != nil {
		log.Fatal("Error unmarshaling file: ", err)
	}
}
