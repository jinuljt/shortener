package shortener

import (
	"fmt"
	"testing"
)

var StartBase = 100000000

func TestEncode(t *testing.T) {
	for i := 0; i < 10; i++ {
		_i := i + StartBase
		fmt.Println(_i, Encode(int64(_i)))
	}
}

func TestAll(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		s := Encode(int64(i))
		_i := Decode(s)
		if int64(i) != _i {
			t.Fatalf("%d encoding to %s\n %s decoding to %d", i, s, s, _i)
		}
	}
}
