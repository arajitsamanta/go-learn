package interfaces

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/tour/reader"
)

/*
GoReaders - The io package specifies the io.Reader interface, which represents the read end of a stream of data.

The Go standard library contains many implementations of these interfaces, including files, network connections, compressors, ciphers, and others.

The io.Reader interface has a Read method:

	func (T) Read(b []byte) (n int, err error)

Read populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an io.EOF error when the stream ends.

The example code creates a strings.Reader and consumes its output 8 bytes at a time
*/
func GoReaders() {
	fmt.Printf("****Running interfaces.GoReaders(), Go readers. \n")
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

type myReader struct{}

func (r myReader) Read(s []byte) (int, error) {
	for i := range s {
		s[i] = 'A'
	}

	return len(s), nil
}

//ReaderExcercise - Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
func ReaderExcercise() {
	fmt.Printf("****Running interfaces.ReaderExcercise(), Go reader excercise. \n")
	reader.Validate(myReader{})
}

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	l, e := r.r.Read(b)
	for i, c := range b {
		if c <= 'Z' && c >= 'A' {
			b[i] = (c-'A'+13)%26 + 'A'
		} else if c >= 'a' && c <= 'z' {
			b[i] = (c-'a'+13)%26 + 'a'
		}
	}
	return l, e
}

/*
ROT13Reader - A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.

For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).

Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13(https://en.wikipedia.org/wiki/ROT13)
substitution cipher to all alphabetical characters.

The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.
*/
func ROT13Reader() {
	fmt.Printf("****Running interfaces.ROT13Reader(), Go reader excercise - Rotate by 13\n")
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
