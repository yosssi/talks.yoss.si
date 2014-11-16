// +build ignore
package main

import "os"

// START OMIT

var exit = os.Exit // HL

func main() {
	if len(os.Args) < 2 {
		exit(1) // HL
	}
}

// END OMIT
