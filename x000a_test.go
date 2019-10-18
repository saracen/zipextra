package zipextra

import (
	"encoding/hex"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNTFS(t *testing.T) {
	tests := []struct {
		attributes []NTFSAttribute
		raw        string
	}{
		{
			[]NTFSAttribute{
				NTFSTimeAttribute{
					time.Date(2018, time.January, 1, 2, 3, 4, 0, time.UTC),
					time.Date(2018, time.January, 3, 0, 18, 10, 0, time.UTC),
					time.Date(2018, time.January, 3, 0, 23, 58, 0, time.UTC),
				},
			},
			"0a0020000000000001001800007caca8a482d30100cdfb552884d301006368252984d301",
		},
		{
			[]NTFSAttribute{
				NTFSTimeAttribute{
					time.Date(2019, time.January, 1, 4, 7, 10, 0, time.UTC),
					time.Date(2019, time.January, 2, 5, 8, 11, 0, time.UTC),
					time.Date(2019, time.January, 3, 6, 9, 12, 0, time.UTC),
				},
				NTFSRawAttribute{
					RawTag:  0xffff,
					RawData: []byte("foobar"),
				},
			},
			"0a002a00000000000100180000ab9c7787a1d40180af262859a2d40100b4b0d82aa3d401ffff0600666f6f626172",
		},
	}

	for _, test := range tests {
		// encode
		raw := NewNTFS(test.attributes...).Encode()
		require.Equal(t, test.raw, hex.EncodeToString(raw))

		// decode
		ntfs, err := testHeader(t, raw, ExtraFieldNTFS).NTFS()
		require.NoError(t, err)
		assert.Equal(t, test.raw, hex.EncodeToString(ntfs.Encode()))
	}
}
