package zipextra

import (
	"encoding/hex"
	"testing"
	"time"
)

func TestExtendedTimestamp(t *testing.T) {
	tests := []struct {
		modTime time.Time
		raw     string
	}{
		{time.Time{}, "55540000"},
		{time.Date(2010, time.January, 1, 1, 2, 3, 4, time.UTC), "55540500018b493d4b"},
	}

	for _, test := range tests {
		// encode
		raw := NewExtendedTimestamp(test.modTime).Encode()
		if test.raw != hex.EncodeToString(raw) {
			t.Errorf("expected %s, got %s", test.raw, hex.EncodeToString(raw))
		}

		// decode
		etime, err := testHeader(t, raw, ExtraFieldExtTime).ExtendedTimestamp()
		if err != nil {
			t.Fatal(err)
		}

		if test.raw != hex.EncodeToString(etime.Encode()) {
			t.Errorf("expected %s, got %s", test.raw, hex.EncodeToString(etime.Encode()))
		}
	}
}
