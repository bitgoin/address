/*
 * Copyright (c) 2016, Shinya Yagyu
 * All rights reserved.
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 * 1. Redistributions of source code must retain the above copyright notice,
 *    this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright notice,
 *    this list of conditions and the following disclaimer in the documentation
 *    and/or other materials provided with the distribution.
 * 3. Neither the name of the copyright holder nor the names of its
 *    contributors may be used to endorse or promote products derived from this
 *    software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
 * LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 * SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 * CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 */

package btcaddr

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"

	"github.com/monarj/wallet/params"
	"github.com/utamaro/btcaddr/base58"
	"github.com/utamaro/btcaddr/btcec"
	"golang.org/x/crypto/ripemd160"
)

var (
	//BitcoinMain is params for main net.
	BitcoinMain = &Params{
		DumpedPrivateKeyHeader: []byte{128},
		AddressHeader:          0,
		HDPrivateKeyID:         []byte{0x04, 0x88, 0xad, 0xe4},
		HDPublicKeyID:          []byte{0x04, 0x88, 0xb2, 0x1e},
	}
	//BitcoinTest is params for test net.
	BitcoinTest = &Params{
		DumpedPrivateKeyHeader: []byte{239},
		AddressHeader:          111,
		HDPrivateKeyID:         []byte{0x04, 0x35, 0x83, 0x94},
		HDPublicKeyID:          []byte{0x04, 0x35, 0x87, 0xcf},
	}
)

//Params is parameters of the coin.
type Params struct {
	DumpedPrivateKeyHeader []byte
	AddressHeader          byte
	HDPrivateKeyID         []byte
	HDPublicKeyID          []byte
}

//PublicKey represents public key for bitcoin
type PublicKey struct {
	*btcec.PublicKey
	isCompressed bool
	param        *Params
}

//PrivateKey represents private key for bitcoin
type PrivateKey struct {
	*btcec.PrivateKey
	PublicKey *PublicKey
}

//NewPublicKey returns PublicKey struct using public key hex string.
func NewPublicKey(pubKeyByte []byte, param *Params) (*PublicKey, error) {
	secp256k1 := btcec.S256()
	key, err := btcec.ParsePubKey(pubKeyByte, secp256k1)
	if err != nil {
		return nil, err
	}
	isCompressed := false
	if len(pubKeyByte) == btcec.PubKeyBytesLenCompressed {
		isCompressed = true
	}
	return &PublicKey{
		PublicKey:    key,
		isCompressed: isCompressed,
		param:        param,
	}, nil
}

//FromWIF gets PublicKey and PrivateKey from private key of WIF format.
func FromWIF(wif string, param *Params) (*PrivateKey, error) {
	secp256k1 := btcec.S256()
	pb, err := base58.Decode(wif)
	if err != nil {
		return nil, err
	}
	ok := false
	for _, h := range param.DumpedPrivateKeyHeader {
		if pb[0] == h {
			ok = true
		}
	}
	if !ok {
		return nil, errors.New("wif is invalid")
	}
	isCompressed := false
	if len(pb) == btcec.PrivKeyBytesLen+2 && pb[btcec.PrivKeyBytesLen+1] == 0x01 {
		pb = pb[:len(pb)-1]
		isCompressed = true
		log.Println("compressed")
	}

	//Get the raw public
	priv, pub := btcec.PrivKeyFromBytes(secp256k1, pb[1:])
	return &PrivateKey{
		PrivateKey: priv,
		PublicKey: &PublicKey{
			PublicKey:    pub,
			isCompressed: isCompressed,
			param:        param,
		},
	}, nil
}

//NewPrivateKey creates and returns PrivateKey from bytes.
func NewPrivateKey(pb []byte, param *Params) *PrivateKey {
	secp256k1 := btcec.S256()
	priv, pub := btcec.PrivKeyFromBytes(secp256k1, pb)
	return &PrivateKey{
		PrivateKey: priv,
		PublicKey: &PublicKey{
			PublicKey:    pub,
			isCompressed: true,
			param:        param,
		},
	}
}

//Generate generates random PublicKey and PrivateKey.
func Generate(param *Params) (*PrivateKey, error) {
	secp256k1 := btcec.S256()
	prikey, err := btcec.NewPrivateKey(secp256k1)
	if err != nil {
		return nil, err
	}
	key := &PrivateKey{
		PublicKey: &PublicKey{
			PublicKey:    prikey.PubKey(),
			isCompressed: true,
			param:        param,
		},
		PrivateKey: prikey,
	}

	return key, nil
}

//Sign sign data.
func (priv *PrivateKey) Sign(hash []byte) ([]byte, error) {
	sig, err := priv.PrivateKey.Sign(hash)
	if err != nil {
		return nil, err
	}
	return sig.Serialize(), nil
}

//WIFAddress returns WIF format string from PrivateKey
func (priv *PrivateKey) WIFAddress() string {
	p := priv.Serialize()
	if priv.PublicKey.isCompressed {
		p = append(p, 0x1)
	}
	p = append(p, 0x0)
	copy(p[1:], p[:len(p)-1])
	p[0] = priv.PublicKey.param.DumpedPrivateKeyHeader[0]
	return base58.Encode(p)
}

//Serialize serializes public key depending on isCompressed.
func (pub *PublicKey) Serialize() []byte {
	if pub.isCompressed {
		return pub.SerializeCompressed()
	}
	return pub.SerializeUncompressed()
}

//AddressBytes returns bitcoin address  bytes from PublicKey
func (pub *PublicKey) AddressBytes() []byte {
	//Next we get a sha256 hash of the public key generated
	//via ECDSA, and then get a ripemd160 hash of the sha256 hash.
	shadPublicKeyBytes := sha256.Sum256(pub.Serialize())

	ripeHash := ripemd160.New()
	if _, err := ripeHash.Write(shadPublicKeyBytes[:]); err != nil {
		log.Fatal(err)
	}
	ripeHashedBytes := ripeHash.Sum(nil)
	ripeHashedBytes = append(ripeHashedBytes, 0x0)
	copy(ripeHashedBytes[1:], ripeHashedBytes[:len(ripeHashedBytes)-1])
	ripeHashedBytes[0] = params.AddressHeader

	return ripeHashedBytes[1:]
}

//Address returns bitcoin address from PublicKey
func (pub *PublicKey) Address() string {
	return base58.Encode(pub.AddressBytes())
}

//DecodeAddress converts bitcoin address to hex form.
func DecodeAddress(addr string) ([]byte, error) {
	pb, err := base58.Decode(addr)
	if err != nil {
		return nil, err
	}
	return pb[1:], nil
}

//Verify verifies signature is valid or not.
func (pub *PublicKey) Verify(signature []byte, data []byte) error {
	secp256k1 := btcec.S256()
	sig, err := btcec.ParseSignature(signature, secp256k1)
	if err != nil {
		return err
	}
	valid := sig.Verify(data, pub.PublicKey)
	if !valid {
		return fmt.Errorf("signature is invalid")
	}
	return nil
}
