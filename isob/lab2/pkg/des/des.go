package des

func Encode(text, key string) string {
	encoded := ""
	offset := len(key)
	for _, v := range text {
		encoded += string(int(v) + offset)
	}
	return encoded
}

func Decode(text, key string) string {
	decoded := ""
	offset := len(key)
	for _, v := range text {
		decoded += string(int(v) - offset)
	}
	return decoded
}
