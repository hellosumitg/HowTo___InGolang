package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	payload := []byte("Hello! Golang developers")
	hashAndBroadcast(newHashReader(payload))
}

// `HashReader` interface represents a reader that can compute a hash.
type HashReader interface {
	io.Reader
	hash() string
}

// `hashReader` is an implementation of the HashReader interface.
type hashReader struct {
	*bytes.Reader
	buf *bytes.Buffer
}

// `newHashReader()` creates a new hashReader instance with the given byte slice.
func newHashReader(b []byte) *hashReader {
	return &hashReader{
		Reader: bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}

// `hash()` computes the hash of the underlying byte slice and returns it as a string.
func (h *hashReader) hash() string {
	return hex.EncodeToString(h.buf.Bytes())
}

// `Older Way (i.e without Interface)` which just converts the `string` to `bytes` but didn't convert the produced `bytes` into `string`
// func hashAndBroadcast(r io.Reader) error {
// 	b, err := io.ReadAll(r)
// 	if err != nil {
// 		return err
// 	}

// 	hash := sha1.Sum(b)
// 	fmt.Println(hex.EncodeToString(hash[:]))

// 	return broadcast(r)
// }

// `hashAndBroadcast()` reads from the HashReader, computes the hash, prints it,
// and then broadcasts the contents to the `broadcast()` function.
func hashAndBroadcast(r HashReader) error {
	// hash := r.(*hashReader).hash()
	hash := r.hash() // Compute the hash using the hashReader implementation
	fmt.Println("Hash or bytes form of the sentence or string: ", hash)

	return broadcast(r) // Broadcast the contents to the broadcast function
}

// `broadcast()` reads from the io.Reader and prints the string representation of the bytes.
func broadcast(r io.Reader) error {
	b, err := io.ReadAll(r) // Read all the data from the reader
	if err != nil {
		return err
	}

	fmt.Println("String of the bytes: ", string(b))

	return nil
}
