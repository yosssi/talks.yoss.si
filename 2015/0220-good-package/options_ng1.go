// +build ignore
package main

// START OMIT

type Dog struct {
	Name string
}

func NewDog(name string) *Dog {
	d := &Dog{Name: name}
	// Some initialization processes go here.
	return d
}

func main() {
	d := NewDog("Taro")
}

// END OMIT
