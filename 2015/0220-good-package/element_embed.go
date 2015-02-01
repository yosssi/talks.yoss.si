// +build ignore
package main

// START OMIT
// elementBase has fields and methods which are common among elements. // HL
type elementBase struct { // HL
	children []element // HL
} // HL

func (e *elementBase) Append(child element) { // HL
	e.children = append(e.children, child) // HL
} // HL

type htmlTag struct {
	elementBase // embed the elementBase struct into the htmlTag struct. // HL
}

type include struct {
	elementBase // embed the `elementBase` struct into the include struct. // HL
}

// END OMIT
