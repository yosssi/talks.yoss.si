// +build ignore
package main

import "io"

// START OMIT

// element represents a template file's element.
type element struct {
	eType string // eType represents a type of the element. // HL
}

// WriteTo writes the element's content to the writer.
func (e *element) WriteTo(w io.Writer) (int64, error) {
	switch e.eType { // HL
	case "HTMLTag": // HL
		// Write its HTML content. // HL
	case "IncludeOperation": // HL
		// Load the other template and write the content. // HL
	case "JSHelper": // HL
		// Write its JavaScript content. // HL
	case "xxx": // HL
		// ... // HL
	} // HL
}

// END OMIT
