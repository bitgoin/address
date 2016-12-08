//Copyright (c) 2012 Tommi Virtanen

//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in
//all copies or substantial portions of the Software.

//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
//THE SOFTWARE.

package base58

import (
	"math/big"
	"testing"
)

type testpair struct {
	decoded int64
	encoded string
}

var pairs = []testpair{
	{10002343, "tGMC"},
	{1000, "JF"},
	{0, ""},
}

func TestEncode(t *testing.T) {
	for _, p := range pairs {
		buf := []byte("noise")
		buf = EncodeBig(buf, big.NewInt(p.decoded))
		if string(buf) != "noise"+p.encoded {
			t.Errorf("unexpected result: %q != %q", string(buf), p.encoded)
		}
	}
}

func TestDecode(t *testing.T) {
	for _, data := range pairs {
		buf := []byte(data.encoded)
		n, err := DecodeToBig(buf)
		if err != nil {
			t.Errorf("decoding %q failed: %v", data.encoded, err)
		}
		if n.Int64() != data.decoded {
			t.Errorf("unexpected result: %v != %v", n, data.decoded)
		}
	}
}

func TestDecodeCorrupt(t *testing.T) {
	type corrupt struct {
		input  string
		offset int
	}
	examples := []corrupt{
		{"!!!!", 0},
		{"x===", 1},
		{"x0", 1},
		{"xl", 1},
		{"xI", 1},
		{"xO", 1},
	}

	for _, e := range examples {
		_, err := DecodeToBig([]byte(e.input))
		switch err := err.(type) {
		case CorruptInputError:
			if int(err) != e.offset {
				t.Errorf("Corruption in %q at offset %v, want %v", e.input, int(err), e.offset)
			}
		default:
			t.Error("Decoder failed to detect corruption in", e)
		}
	}
}
