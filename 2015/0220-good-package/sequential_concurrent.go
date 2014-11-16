// +build ignore
package main

import "fmt"

// START OMIT

func main() {
	// Call the `Compile` function sequentially.
	resultPath1 := Compile("file1")
	resultPath2 := Compile("file2")

	// Calling it concurrently is also easy.
	originalFilePaths := []string{"file3", "file4", "file5"}
	chResultPath := make(chan string)

	for _, path := range originalFilePaths {
		go func(path string) {
			chResultPath <- Compile(path)
		}(path)
	}

	for resultPath := range chResultPath {
		fmt.Println(resultPath)
	}
}

// END OMIT
