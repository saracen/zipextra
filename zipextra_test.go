package zipextra

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testHeader(t *testing.T, raw []byte, identifier uint16) ExtraField {
	b := NewBuffer(raw)
	size := len(b.Bytes()) - 4
	i := b.Read16()
	s := int(b.Read16())

	require.Equal(t, identifier, i)
	require.Equal(t, size, s)

	return ExtraField(b.Bytes())
}

func TestParse(t *testing.T) {
	extraFields := [][]byte{
		NewInfoZIPNewUnix(big.NewInt(1), big.NewInt(1)).Encode(),
		NewInfoZIPUnicodeComment("foobar").Encode(),
	}

	var extra []byte
	for _, field := range extraFields {
		extra = append(extra, field...)
	}

	fields, err := Parse(extra)
	require.NoError(t, err)

	assert.Contains(t, fields, ExtraFieldUnixN)
	assert.Contains(t, fields, ExtraFieldUCom)
}
