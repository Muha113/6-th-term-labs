package main

import (
	"fmt"

	"github.com/Muha113/6-th-term-labs/isob/lab1/caesar"
	"github.com/Muha113/6-th-term-labs/isob/lab1/model"
	"github.com/Muha113/6-th-term-labs/isob/lab1/vigenere"
)

func main() {
	fmt.Println("+++++Caesar+++++")

	c := model.CaesarJSON{}

	fmt.Println("-----Test case 1-----")

	c.Unmarshal("caesar1.json")
	fmt.Printf("Text: |%s|\n", c.Text)
	encoded := caesar.Encode(c.Text, c.Key)
	fmt.Printf("Encoded: |%s|\n", encoded)
	decoded := caesar.Decode(encoded, c.Key)
	fmt.Printf("Decoded: |%s|\n", decoded)

	fmt.Println("-----Test case 2-----")

	c.Unmarshal("caesar2.json")
	fmt.Printf("Text: |%s|\n", c.Text)
	encoded = caesar.Encode(c.Text, c.Key)
	fmt.Printf("|%s|\n", encoded)
	decoded = caesar.Decode(encoded, c.Key)
	fmt.Printf("|%s|\n", decoded)

	fmt.Println("+++++Vigenere+++++")

	v := model.VigenereJSON{}

	fmt.Println("-----Test case 1-----")

	v.Unmarshal("vigenere1.json")
	fmt.Printf("Text: |%s|\n", v.Text)
	encoded = vigenere.Encode(v.Text, v.Key)
	fmt.Printf("Encoded: |%s|\n", encoded)
	decoded = vigenere.Decode(encoded, v.Key)
	fmt.Printf("Decoded: |%s|\n", decoded)

	fmt.Println("-----Test case 2-----")

	v.Unmarshal("vigenere2.json")
	fmt.Printf("Text: |%s|\n", v.Text)
	encoded = vigenere.Encode(v.Text, v.Key)
	fmt.Printf("Encoded: |%s|\n", encoded)
	decoded = vigenere.Decode(encoded, v.Key)
	fmt.Printf("Decoded: |%s|\n", decoded)
}
