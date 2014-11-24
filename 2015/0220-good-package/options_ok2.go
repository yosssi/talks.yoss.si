// +build ignore
package main

// START OMIT

type Dog struct {
	name string
	age  int // HL
}

type Options struct {
	Name string
	Age  int // HL
}

func NewDog(opts *Options) *Dog {
	if opts == nil {
		opts = &Options{}
	}
	d := &Dog{name: opts.Name, age: opts.Age} // HL
	return d // Some initialization might be done before returning `d`.
}

func main() {
	d := NewDog(&Options{Name: "Taro", Age: 7}) // HL
}

// END OMIT
