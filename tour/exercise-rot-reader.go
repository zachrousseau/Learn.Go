// Exercise: rot13Reader
// A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.

// For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).

// Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.

// The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, err error){ // points to r. r is a rot13Reader struct with the attribute R and a function Read

	n, err = r.r.Read(b) // Read from io.Reader. n is the bytes that were read. It's really just a pointer in memory 

	for x := range b {
		
		// Capital Letters
		if b[x] > 64 && b[x] < 91{
			b[x] = b[x] + 13
			
			if b[x] > 90{ // Wrap around
				b[x] = b[x] - 26
			}


		// Lowercase letters 
		} else if b[x] > 96 && b[x] < 123 {
			b[x] = b[x] + 13
			
			if b[x] > 122{ // wrap around 
				b[x] = b[x] - 26
			}
		}
	}

	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!") // Readers regular
	r := rot13Reader{s} // Makes a rot13 reader 
	io.Copy(os.Stdout, &r)
}
