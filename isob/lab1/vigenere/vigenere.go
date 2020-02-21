package vigenere

import (
	"unicode/utf8"
)

const (
	LEFT_BOUND  = 32
	RIGHT_BOUND = 122
)

func Encode(text, key string) string {
	encoded := ""
	i := 0
	module := (RIGHT_BOUND - LEFT_BOUND + 1)
	for _, v := range text {
		if i > utf8.RuneCountInString(key)-1 {
			i = 0
		}

		if int(v) < LEFT_BOUND || int(v) > RIGHT_BOUND {
			encoded += string(v)
		} else {
			val := (((int(v) - LEFT_BOUND) + (int(key[i]) - LEFT_BOUND)) % module) + LEFT_BOUND
			encoded += string(val)
		}
		i++
	}
	return encoded
}

func Decode(text, key string) string {
	decoded := ""
	i := 0
	module := (RIGHT_BOUND - LEFT_BOUND + 1)
	for _, v := range text {
		if i > utf8.RuneCountInString(key)-1 {
			i = 0
		}

		if int(v) < LEFT_BOUND || int(v) > RIGHT_BOUND {
			decoded += string(v)
		} else {
			val := (int(v) - LEFT_BOUND) - (int(key[i]) - LEFT_BOUND)
			if val < 0 {
				val += module
			}
			decoded += string(val + LEFT_BOUND)
		}
		i++
	}
	return decoded
}
