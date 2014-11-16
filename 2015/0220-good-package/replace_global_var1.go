// +build ignore
package main

import "os"

// START OMIT

func main() {
	if len(os.Args) < 2 {
		os.Exit(1) // Can not test this block because os.Exit terminates the process. // HL
	}
}

// END OMIT
