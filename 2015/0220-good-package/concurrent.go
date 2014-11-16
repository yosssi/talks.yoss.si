// +build ignore
package main

// START OMIT

func Compile(path string) <-chan string {
	chPath := make(chan string)

	go func() { // HL
		// Read the simplified CSS format file specified by the paramter, // HL
		// compile it into a pure CSS format data and // HL
		// write the result to a new file. // HL

		// Send the result file's path to the chPath channel. // HL
		chPath <- resultFilePath // HL
	}() // HL

	return chPath
}

// END OMIT
