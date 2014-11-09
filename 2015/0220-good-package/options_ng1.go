// +build ignore
package main

// START OMIT

type Dog struct {
	name string
}

func NewDog(name string) *Dog {
	d := &Dog{name: name}
	return d // Some initialization might be done before returning `d`.
}

func main() {
	d := NewDog("Taro")
}

// END OMIT
