// +build ignore
package main

import "io"

// START OMIT

func Compile(src io.Reader) string { // HL
	// Read the simplified CSS format data from the reader, // HL
	// compile it into a pure CSS format data and
	// write the result to a new file.

	// Return the result file's path.
	return resultFilePath
}

// END OMIT
