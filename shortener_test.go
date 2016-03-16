package shortener

import (
	"fmt"
	"testing"
)

var StartBase = 100000000

func TestEncode(t *testing.T) {
	shorten := NewShortener("")
	for i := 0; i < 10; i++ {
		_i := i + StartBase
		fmt.Println(_i, shorten.Encode(int64(_i)))
	}
}

func TestAll(t *testing.T) {
	shorten := NewShortener("")
	for i := 0; i < 100000; i++ {
		s := shorten.Encode(int64(i))
		_i := shorten.Decode(s)
		if int64(i) != _i {
			t.Fatalf("%d encoding to %s\n %s decoding to %d", i, s, s, _i)
		}
	}
}
