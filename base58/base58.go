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

// Package base58 implements a human-friendly base58 encoding.
//
// As opposed to base64 and friends, base58 is typically used to
// convert integers. You can use big.Int.SetBytes to convert arbitrary
// bytes to an integer first, and big.Int.Bytes the other way around.
package base58

import (
	"math/big"
	"strconv"
)

const alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

var decodeMap [256]byte

func init() {
	for i := 0; i < len(decodeMap); i++ {
		decodeMap[i] = 0xFF
	}
	for i := 0; i < len(alphabet); i++ {
		decodeMap[alphabet[i]] = byte(i)
	}
}

//CorruptInputError representds input is corrupted.
type CorruptInputError int64

//Error is to implement Error interface.
func (e CorruptInputError) Error() string {
	return "illegal base58 data at input byte " + strconv.FormatInt(int64(e), 10)
}

// DecodeToBig a big integer from the bytes. Returns an error on corrupt
// input.
func DecodeToBig(src []byte) (*big.Int, error) {
	n := new(big.Int)
	radix := big.NewInt(58)
	for i := 0; i < len(src); i++ {
		b := decodeMap[src[i]]
		if b == 0xFF {
			return nil, CorruptInputError(i)
		}
		n.Mul(n, radix)
		n.Add(n, big.NewInt(int64(b)))
	}
	return n, nil
}

// EncodeBig encodes src, appending to dst. Be sure to use the returned
// new value of dst.
func EncodeBig(dst []byte, src *big.Int) []byte {
	start := len(dst)
	n := new(big.Int)
	n.Set(src)
	radix := big.NewInt(58)
	zero := big.NewInt(0)

	for n.Cmp(zero) > 0 {
		mod := new(big.Int)
		n.DivMod(n, radix, mod)
		dst = append(dst, alphabet[mod.Int64()])
	}

	for i, j := start, len(dst)-1; i < j; i, j = i+1, j-1 {
		dst[i], dst[j] = dst[j], dst[i]
	}
	return dst
}
