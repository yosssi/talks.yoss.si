// +build ignore
package main

import "os"

// START OMIT

var exit = os.Exit // Define a global variable. // HL

func main() {
	if len(os.Args) < 2 {
		exit(1) // Can test this block by replacing the value of the exit function. // HL
	}
}

// END OMIT
