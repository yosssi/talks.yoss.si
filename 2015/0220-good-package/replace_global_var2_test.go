// +build ignore
package main

import "testing"

// START OMIT

func Test_main(t *testing.T) {
	defer func(orig func(int)) {
		exit = orig // Restore the value of the exit global variable. // HL
	}(exit)

	exit = func(_ int) {} // Replace the value of the exit function with another one. // HL

	main() // Execute the test.
}

// END OMIT
