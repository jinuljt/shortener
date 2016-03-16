package shortener

import (
	"math"
)

var ALPHABET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func Encode(n int64) string {
	if n == 0 {
		return string(ALPHABET[0])
	}

	s := ""
	base := int64(len(ALPHABET))
	for {
		if n <= 0 {
			return s
		}

		remainder := n % base
		s = string(ALPHABET[remainder]) + s
		n /= base
	}
}

func Decode(s string) int64 {
	var n int64
	n = 0

	alphabetMap := make(map[rune]int64)
	for idx, char := range ALPHABET {
		alphabetMap[char] = int64(idx)
	}

	base := float64(len(ALPHABET))
	for idx, char := range reverse(s) {
		n += alphabetMap[char] * int64(math.Pow(base, float64(idx)))
	}
	return n
}
