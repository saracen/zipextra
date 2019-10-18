package zipextra

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInfoZIPUnicodeComment(t *testing.T) {
	tests := []struct {
		version uint8
		crc32   uint32
		comment string
		raw     string
	}{
		{1, 0x3c27b58d, "a comment 🌩", "75631300018db5273c6120636f6d6d656e7420f09f8ca9"},
		{1, 0x9cc1778e, "another comment ¿", "75631700018e77c19c616e6f7468657220636f6d6d656e7420c2bf"},
		{1, 0x643c7f60, "wild comment appeared! ™", "75631f0001607f3c6477696c6420636f6d6d656e742061707065617265642120e284a2"},
	}

	for _, test := range tests {
		// encode
		raw := NewInfoZIPUnicodeComment(test.comment).Encode()
		require.Equal(t, test.raw, hex.EncodeToString(raw))

		// decode
		ucom, err := testHeader(t, raw, ExtraFieldUCom).InfoZIPUnicodeComment()
		require.NoError(t, err)
		assert.Equal(t, test.raw, hex.EncodeToString(ucom.Encode()))
	}
}
