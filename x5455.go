package zipextra

import "time"

// ExtraFieldExtTime is the Extended Timestamp Extra Field
// identifier.
const ExtraFieldExtTime uint16 = 0x5455

// ExtendedTime is the Extended Timestamp Extra Field structure for holding a
// file entry's last modification time.
type ExtendedTimestamp struct {
	ModTime time.Time
}

// NewExtendedTimestamp returns a new ExtendedTime extra field structure.
func NewExtendedTimestamp(modTime time.Time) ExtendedTimestamp {
	return ExtendedTimestamp{
		ModTime: modTime,
	}
}

// Encode encodes the ExtendedTimestamp extra field.
func (field ExtendedTimestamp) Encode() []byte {
	buf := NewBuffer([]byte{})
	defer buf.WriteHeader(ExtraFieldExtTime)()

	if field.ModTime.IsZero() {
		return buf.Bytes()
	}

	buf.Write8(1)
	buf.Write32(uint32(field.ModTime.Unix()))

	return buf.Bytes()
}

// ExtendedTime returns the decoded ExtendedTime extra field.
func (ef ExtraField) ExtendedTimestamp() (field ExtendedTimestamp, err error) {
	buf := NewBuffer(ef)
	if buf.Available() == 0 {
		return field, nil
	}

	if buf.Read8()&1 == 1 {
		if buf.Available() < 4 {
			return field, ErrInvalidExtraFieldFormat
		}

		field.ModTime = time.Unix(int64(buf.Read32()), 0)
	}

	return
}
