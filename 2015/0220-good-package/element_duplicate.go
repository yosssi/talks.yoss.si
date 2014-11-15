// +build ignore
package main

import "io"

// START OMIT

type element interface {
	io.WriterTo
	Append(child element) // HL
}

type htmlTag struct {
	children []element // children is a common field among elements. // HL
}

func (h *htmlTag) Append(child element) { // Append is a common method among elements. // HL
	h.children = append(h.children, child) // HL
} // HL

type include struct {
	children []element // HL
}

func (i *include) Append(child element) { // HL
	i.children = append(i.children, child) // HL
} // HL

// END OMIT
