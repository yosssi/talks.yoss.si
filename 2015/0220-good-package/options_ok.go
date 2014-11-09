// +build ignore
package main

// START OMIT

type Dog struct {
	name string
}

type Options struct { // HL
	Name string // HL
}

func NewDog(opts *Options) *Dog { // HL
	if opts == nil { // HL
		opts := &Options{} // HL
	}

	d := &Dog{name: opts.Name} // HL

	return d // HL
} // HL

func main() {
	d := NewDog(&Options{Name: "Taro"}) // HL
}

// END OMIT
