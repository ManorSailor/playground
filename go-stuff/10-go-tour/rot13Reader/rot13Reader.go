// Exercise: Rot13Reader
// Link: https://go.dev/tour/methods/23

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

const ROT = 13

func (r rot13Reader) Read(processed []byte) (int, error) {
	raw := r.r

	n, err := raw.Read(processed)

	for i := range n {
		processed[i] = rottify(processed[i])
	}

	return n, err
}

func rottify(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		index := b - 'A'
		index = (index + ROT) % 26
		return index + 'A'
	}

	if b >= 'a' && b <= 'z' {
		index := b - 'a'
		index = (index + ROT) % 26
		return index + 'a'
	}

	return b
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
