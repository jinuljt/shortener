package shortener

import (
	"math"
)

type Shortener struct {
	alphabet string
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func NewShortener(alphabet string) *Shortener {
	if alphabet == "" {
		alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	}
	return &Shortener{
		alphabet: alphabet,
	}
}

func (shorten *Shortener) Encode(n int64) string {
	if n == 0 {
		return string(shorten.alphabet[0])
	}

	s := ""
	base := int64(len(shorten.alphabet))
	for {
		if n <= 0 {
			return s
		}

		remainder := n % base
		s = string(shorten.alphabet[remainder]) + s
		n /= base
	}
}

func (shorten *Shortener) Decode(s string) int64 {
	var n int64
	n = 0

	alphabetMap := make(map[rune]int64)
	for idx, char := range shorten.alphabet {
		alphabetMap[char] = int64(idx)
	}

	base := float64(len(shorten.alphabet))
	for idx, char := range reverse(s) {
		n += alphabetMap[char] * int64(math.Pow(base, float64(idx)))
	}
	return n
}
