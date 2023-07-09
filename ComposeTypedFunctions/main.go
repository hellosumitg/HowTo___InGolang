package main

import (
	"fmt"
)

// Uncomment for 1st time 1st chance code execution
// type Server struct{}

// Uncomment for 1st time 2nd chance as well as 2nd time 1st and 2nd chance both code execution
// Server holds the filename transformation function.
type Server struct {
	filenameTransformFunc TransformFunc
}

// Uncomment 1st time 2nd chance as well as 2nd time 1st and 2nd chance both code execution
// TransformFunc represents a function that transforms a string.
type TransformFunc func(string) string

// handleRequest processes the filename using the provided transformation function.
func (s *Server) handleRequest(filename string) error {
	// hash := sha256.Sum256([]byte(filename))    // Uncomment for 1st time 1st chance code execution
	// newFilename := hex.EncodeToString(hash[:]) // Uncomment for 1st time 1st chance code execution
	// newFilename := hashFilename(filename) // Uncomment for 1st time 2nd chance code execution
	newFilename := s.filenameTransformFunc(filename) // Uncomment for 2nd time 1st and 2nd chance both code execution
	fmt.Println("New Filename: ", newFilename)
	return nil
}

/*
What if we want to use a different hashing algorithm like SHA1 or something else?
Or what if we want to prefix the fileName with G_? Additionally,
what if we want to use HMAC for hashing? And what if we want to make these changes dynamically?
To address these requirements, function or typed function composability comes into play.
We can use an interface for this purpose, which is a good option.
However, if the desired functionality does not require any state, using an interface may not make sense.
In such cases, instead of implementing the interface with a struct, we can use a function type, as it doesn't hold any state.
So, if we are certain that our interface won't have state, we can make it a typed function.
*/

// Use only for 1st time 2nd chance execution
// hashFilename hashes the filename using SHA256.
// func hashFilename(filename string) string {
// 	hash := sha256.Sum256([]byte(filename))
// 	newFilename := hex.EncodeToString(hash[:])
// 	return newFilename
// }

// prefixFilename prefixes the filename with a given prefix.

// Uncomment for 2nd time 1st chance code execution
// func GPrefixFilename(fileName string) string {
// 	return "G_" + fileName
// }

// Uncomment for 2nd time 2nd chance time code execution
func prefixFilename(prefix string) TransformFunc {
	return func(filename string) string {
		return prefix + filename
	}
}

func main() {
	s := &Server{
		// for 1st time code execution keep this block empty
		// filenameTransformFunc: hashFilename, // Uncomment for 1st time 2nd chance code execution
		// filenameTransformFunc: GPrefixFilename, // Uncomment for 2nd time code execution
		filenameTransformFunc: prefixFilename("SKG_"), // Uncomment for 2nd time 2nd chance code execution
	}
	s.handleRequest("cool_picture.jpg")
}
