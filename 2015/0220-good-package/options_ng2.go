// +build ignore
package main

// START OMIT

type Dog struct {
	name string
	age  int // HL
}

func (d *Dog) SetAge(age int) *Dog { // HL
	d.age = age // HL
	return d    // HL
}

func NewDog(name string) *Dog {
	d := &Dog{name: name}
	return d // Some initialization might be done before returning `d`.
}

func main() {
	d := NewDog("Taro").SetAge(7) // HL
}

// END OMIT
