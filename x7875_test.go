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
		{1, big.NewInt(0), big.NewInt(0), "75780300010000"},
		{1, big.NewInt(1), big.NewInt(1), "757805000101010101"},
		{1, big.NewInt(255), big.NewInt(255), "757805000101ff01ff"},
		{1, big.NewInt(1000), big.NewInt(1000), "757807000102e80302e803"},

		// unusual
		{1, big.NewInt(math.MaxInt32), big.NewInt(math.MaxInt32), "75780b000104ffffff7f04ffffff7f"},
		{1, big.NewInt(math.MaxInt64), big.NewInt(math.MaxInt64), "757813000108ffffffffffffff7f08ffffffffffffff7f"},
	}

	for _, test := range tests {
		// encode
		raw := NewInfoZIPNewUnix(test.uid, test.gid).Encode()
		require.Equal(t, test.raw, hex.EncodeToString(raw))

		// decode
		unix, err := testHeader(t, raw, ExtraFieldUnixN).InfoZIPNewUnix()
		require.NoError(t, err)
		assert.Equal(t, test.raw, hex.EncodeToString(unix.Encode()))
	}
}
