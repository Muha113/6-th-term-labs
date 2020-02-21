package caesar

const (
	LEFT_BOUND  = 32
	RIGHT_BOUND = 122
)

func Encode(text string, key int) string {
	encoded := ""
	for _, v := range text {
		if int(v) > RIGHT_BOUND || int(v) < LEFT_BOUND {
			continue
		}

		if int(v)+key <= RIGHT_BOUND {
			encoded += string(int(v) + key)
		} else {
			tmp := RIGHT_BOUND - int(v)
			encoded += string(LEFT_BOUND + (key - tmp - 1))
		}
	}
	return encoded
}

func Decode(text string, key int) string {
	decoded := ""
	for _, v := range text {
		if int(v) > RIGHT_BOUND || int(v) < LEFT_BOUND {
			continue
		}

		if int(v)-key >= LEFT_BOUND {
			decoded += string(int(v) - key)
		} else {
			tmp := int(v) - LEFT_BOUND
			decoded += string(RIGHT_BOUND - (key - tmp - 1))
		}
	}
	return decoded
}
