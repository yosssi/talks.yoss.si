// +build ignore
package main

import "io"

// START OMIT

func Compile(dst io.Writer, src io.Reader) (int, error) { // HL
	// Read the simplified CSS format data from the reader,
	// compile it into a pure CSS format data.

	// Write the result data to the writer. // HL
	return dst.Write(resultData) // HL
}

// END OMIT
