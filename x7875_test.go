package zipextra

import (
	"encoding/hex"
	"math"
	"math/big"
	"testing"
)

func TestInfoZIPNewUnix(t *testing.T) {
	tests := []struct {
		version uint8
		uid     *big.Int
		gid     *big.Int
		raw     string
	}{
		// normal ids
		{1, big.NewInt(0), big.NewInt(0), "75780b000104000000000400000000"},
		{1, big.NewInt(1), big.NewInt(1), "75780b000104010000000401000000"},
		{1, big.NewInt(255), big.NewInt(255), "75780b000104ff00000004ff000000"},
		{1, big.NewInt(1000), big.NewInt(1000), "75780b000104e803000004e8030000"},

		// unusual
		{1, big.NewInt(math.MaxInt32), big.NewInt(math.MaxInt32), "75780b000104ffffff7f04ffffff7f"},
		{1, big.NewInt(math.MaxInt64), big.NewInt(math.MaxInt64), "757813000108ffffffffffffff7f08ffffffffffffff7f"},
		{1, new(big.Int).SetBytes([]byte{0x0a, 0x0b}), big.NewInt(math.MaxInt64), "75780f0001040b0a000008ffffffffffffff7f"},
	}

	for _, test := range tests {
		// encode
		raw := NewInfoZIPNewUnix(test.uid, test.gid).Encode()
		if test.raw != hex.EncodeToString(raw) {
			t.Errorf("expected %s, got %s", test.raw, hex.EncodeToString(raw))
		}

		// decode
		unix, err := testHeader(t, raw, ExtraFieldUnixN).InfoZIPNewUnix()
		if err != nil {
			t.Fatal(err)
		}

		if test.raw != hex.EncodeToString(unix.Encode()) {
			t.Errorf("expected %s, got %s", test.raw, hex.EncodeToString(unix.Encode()))
		}
		if test.uid.String() != unix.Uid.String() {
			t.Errorf("expected uid %s, got %s", test.uid.String(), unix.Uid.String())
		}
		if test.gid.String() != unix.Gid.String() {
			t.Errorf("expected gid %s, got %s", test.gid.String(), unix.Gid.String())
		}
	}
}
