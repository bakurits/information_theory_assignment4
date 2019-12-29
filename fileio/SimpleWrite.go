package fileio

import (
	"io"
	"io/ioutil"
)

// SimpleWrite converts binary to file
// works only for files with length of multiples 8
func SimpleWrite(r io.Reader, w io.Writer) {
	data, _ := ioutil.ReadAll(r)

	for i := 0; i < len(data); i += 8 {
		newByte, err := ByteArrayToByte(data[i : i+8])
		var arr = []byte{newByte}
		if err != nil {
			return
		}
		_, _ = w.Write(arr)
	}
}
