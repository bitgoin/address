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

package address

var (
	//BitcoinMain is params for main net.
	BitcoinMain = &Params{
		DumpedPrivateKeyHeader: []byte{128},
		AddressHeader:          []byte{0},
		P2SHHeader:             []byte{5},
		HDPrivateKeyID:         []byte{0x04, 0x88, 0xad, 0xe4},
		HDPublicKeyID:          []byte{0x04, 0x88, 0xb2, 0x1e},
	}
	//BitcoinTest is params for test net.
	BitcoinTest = &Params{
		DumpedPrivateKeyHeader: []byte{239},
		AddressHeader:          []byte{111},
		P2SHHeader:             []byte{196},
		HDPrivateKeyID:         []byte{0x04, 0x35, 0x83, 0x94},
		HDPublicKeyID:          []byte{0x04, 0x35, 0x87, 0xcf},
	}
	//MonacoinMain is params for monacoin main net.
	MonacoinMain = &Params{
		DumpedPrivateKeyHeader: []byte{178, 176},
		AddressHeader:          []byte{50},
		P2SHHeader:             []byte{5},
		HDPrivateKeyID:         []byte{0x04, 0x88, 0xad, 0xe4},
		HDPublicKeyID:          []byte{0x04, 0x88, 0xb2, 0x1e},
	}
	//LitecoinMain params for Litecoin main net
	LitecoinMain = &Params{
		DumpedPrivateKeyHeader: []byte{0xB0},
		AddressHeader:          []byte{0x30},
		P2SHHeader:             []byte{0x50},
		HDPrivateKeyID:         []byte{0x04, 0x88, 0xad, 0xe4},
		HDPublicKeyID:          []byte{0x04, 0x88, 0xb2, 0x1e},
	}
)
