// +build ignore
package main

// START OMIT

func Compile(path string) <-chan string {
	chPath := make(chan string)

	// Run the compilation processing always concurrently
	// because it involves many I/O blocking processing.
	go func() {
		// Read the simplified CSS format file specified by the paramter,
		// compile it into a pure CSS format data and
		// write the result to a new file.

		// send the result file's path to the chPath channel.
		chPath <- resultFilePath
	}()

	return chPath
}

// END OMIT
