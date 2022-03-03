package zipextra

import (
	"math/big"
	"testing"
)

func testHeader(t *testing.T, raw []byte, identifier uint16) ExtraField {
	b := NewBuffer(raw)
	size := len(b.Bytes()) - 4
	i := b.Read16()
	s := int(b.Read16())

	if identifier != i {
		t.Errorf("expected identifer %d, got %d", identifier, i)
	}
	if size != s {
		t.Errorf("expected size %d, got %d", size, s)
	}

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
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := fields[ExtraFieldUnixN]; !ok {
		t.Error("expected new unix field")
	}
	if _, ok := fields[ExtraFieldUCom]; !ok {
		t.Error("expected comment field")
	}
	if _, ok := fields[ExtraFieldNTFS]; ok {
		t.Error("unexpected ntfs field")
	}
}
