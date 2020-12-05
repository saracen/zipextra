package zipextra

import (
	"encoding/hex"
	"math"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		require.Equal(t, test.raw, hex.EncodeToString(raw))

		// decode
		unix, err := testHeader(t, raw, ExtraFieldUnixN).InfoZIPNewUnix()
		require.NoError(t, err)
		assert.Equal(t, test.raw, hex.EncodeToString(unix.Encode()))
		assert.Equal(t, test.uid.String(), unix.Uid.String())
		assert.Equal(t, test.gid.String(), unix.Gid.String())
	}
}
