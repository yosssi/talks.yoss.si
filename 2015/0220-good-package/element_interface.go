// +build ignore
package main

import "io"

// START OMIT

type element interface { // HL
	io.WriterTo // HL
} // HL

type htmlTag struct{} // HL

func (h *htmlTag) WriteTo(w io.Writer) (int64, error) { // HL
	// Write the html tag's content to the writer. // HL
} // HL

type include struct{} // HL

func (i *include) WriteTo(w io.Writer) (int64, error) { // HL
	// Load the other template and write the content. // HL
} // HL

// END OMIT
