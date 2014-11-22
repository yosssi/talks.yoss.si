// +build ignore
package main

// START OMIT

// privateStruct is used only in this package.
type privateStruct struct{} // HL

// PublicFunc is used not only in this package but also other packages.
func PublicFunc() {}

// END OMIT
