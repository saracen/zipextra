package zipextra_test

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"math/big"

	"github.com/saracen/zipextra"
)

func ExampleZipExtra() {
	// create temporary file
	w, err := ioutil.TempFile("", "zipextra-example")
	if err != nil {
		panic(err)
	}

	// create new zip writer
	zw := zip.NewWriter(w)

	// create a new zip file header
	fh := &zip.FileHeader{Name: "test_file.txt"}

	// add some extra fields
	fh.Extra = append(fh.Extra, zipextra.NewInfoZIPNewUnix(big.NewInt(1000), big.NewInt(1000)).Encode()...)
	fh.Extra = append(fh.Extra, zipextra.NewInfoZIPUnicodeComment("Hello, 世界").Encode()...)

	// create the file
	fw, err := zw.CreateHeader(fh)
	if err != nil {
		panic(err)
	}
	fw.Write([]byte("foobar"))
	zw.Close()

	// open the newly created zip
	zr, err := zip.OpenReader(w.Name())
	if err != nil {
		panic(err)
	}
	defer zr.Close()

	// parse extra fields
	fields, err := zipextra.Parse(zr.File[0].Extra)
	if err != nil {
		panic(err)
	}

	// print extra field information
	for id, field := range fields {
		switch id {
		case zipextra.ExtraFieldUnixN:
			unix, _ := field.InfoZIPNewUnix()
			fmt.Printf("UID: %d, GID: %d\n", unix.Uid, unix.Gid)
		}
	}
	for id, field := range fields {
		switch id {
		case zipextra.ExtraFieldUCom:
			ucom, _ := field.InfoZIPUnicodeComment()
			fmt.Printf("Comment: %s\n", ucom.Comment)
		}
	}
	// Output:
	// UID: 1000, GID: 1000
	// Comment: Hello, 世界
}
