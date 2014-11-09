// +build ignore
package main

// START OMIT

type Dog struct {
	name string
}

func NewDog(name string) *Dog {
	d := &Dog{name: name}
	// Some initialization processes go here.
	return d
}

func main() {
	d := NewDog("Taro")
}

// END OMIT
