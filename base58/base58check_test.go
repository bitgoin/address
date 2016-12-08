//from https://github.com/btcsuite/btcutil/blob/master/base58/base58check_test.go
//under ISC license
//https://github.com/btcsuite/btcutil/blob/master/LICENSE

package base58

import "testing"

var checkEncodingStringTests = []struct {
	in  string
	out string
}{
	{"\x14", "3MNQE1X"},
	{"\x14 ", "B2Kr6dBE"},
	{"\x14-", "B3jv1Aft"},
	{"\x140", "B482yuaX"},
	{"\x141", "B4CmeGAC"},
	{"\x14-1", "mM7eUf6kB"},
	{"\x1411", "mP7BMTDVH"},
	{"\x14abc", "4QiVtDjUdeq"},
	{"\x141234598760", "ZmNb8uQn5zvnUohNCEPP"},
	{"\x14abcdefghijklmnopqrstuvwxyz", "K2RYDcKfupxwXdWhSAxQPCeiULntKm63UXyx5MvEH2"},
	{"\x1400000000000000000000000000000000000000000000000000000000000000", "bi1EWXwJay2udZVxLJozuTb8Meg4W9c6xnmJaRDjg6pri5MBAxb9XwrpQXbtnqEoRV5U2pixnFfwyXC8tRAVC8XxnjK"},
}

func TestBase58Check(t *testing.T) {
	for x, test := range checkEncodingStringTests {
		// test encoding
		if res := Encode([]byte(test.in)); res != test.out {
			t.Errorf("CheckEncode test #%d failed: got %s, want: %s", x, res, test.out)
		}

		// test decoding
		res, err := Decode(test.out)
		if err != nil {
			t.Errorf("CheckDecode test #%d failed with err: %v", x, err)
		} else if string(res) != test.in {
			t.Errorf("CheckDecode test #%d failed: got: %s want: %s", x, res, test.in)
		}
	}

	// test the two decoding failure cases
	// case 1: checksum error
	_, err := Decode("3MNQE1Y")
	if err == nil {
		t.Error("Checkdecode test failed, expected ErrChecksum")
	}
	// case 2: invalid formats (string lengths below 5 mean the version byte and/or the checksum
	// bytes are missing).
	testString := ""
	for len := 0; len < 4; len++ {
		// make a string of length `len`
		_, err = Decode(testString)
		if err == nil {
			t.Error("Checkdecode test failed, expected ErrInvalidFormat")
		}
	}

}
