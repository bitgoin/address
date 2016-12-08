//Copyright (c) 2014 Bryce Neal

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
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
)

//Encode encodes byteData to base58.
func Encode(encoded []byte) string {
	//Perform SHA-256 twice
	hash := sha256.Sum256(encoded)
	hash = sha256.Sum256(hash[:])

	//First 4 bytes if this double-sha'd byte array is the checksum
	//Append this checksum to the input bytes
	encoded = append(encoded, hash[0:4]...)

	//Convert this checksum'd version to a big Int
	bigIntEncodedChecksum := new(big.Int).SetBytes(encoded)

	//Encode the big int checksum'd version into a Base58Checked string
	base58EncodedChecksum := EncodeBig(nil, bigIntEncodedChecksum)

	//Now for each zero byte we counted above we need to prepend a 1 to our
	//base58 encoded string. The rational behind this is that base58 removes 0's (0x00).
	//So bitcoin demands we add leading 0s back on as 1s.
	buffer := make([]byte, 0, len(base58EncodedChecksum))

	//base58 alone is not enough. We need to first count each of the zero bytes
	//which are at the beginning of the encodedCheckSum

	for _, v := range encoded {
		if v != 0 {
			break
		}
		buffer = append(buffer, '1')
	}
	buffer = append(buffer, base58EncodedChecksum...)
	return string(buffer)
}

//Decode decodes base58 value to bytes.
func Decode(value string) ([]byte, error) {
	if len(value) < 5 {
		return nil, errors.New("invalid input")
	}
	publicKeyInt, err := DecodeToBig([]byte(value))
	if err != nil {
		return nil, err
	}

	encodedChecksum := publicKeyInt.Bytes()
	encoded := encodedChecksum[:len(encodedChecksum)-4]
	cksum := encodedChecksum[len(encodedChecksum)-4:]

	buffer := make([]byte, 0, len(encoded))
	for _, v := range value {
		if v != '1' {
			break
		}
		buffer = append(buffer, 0)
	}

	buffer = append(buffer, encoded...)

	//Perform SHA-256 twice
	hash := sha256.Sum256(buffer)
	hash = sha256.Sum256(hash[:])

	if !bytes.Equal(hash[:4], cksum) {
		return nil,
			fmt.Errorf("%s checksum did not match to the embeded one:%s, should be :%s",
				value, hex.EncodeToString(hash[:4]), hex.EncodeToString(cksum))
	}

	return buffer, err
}
