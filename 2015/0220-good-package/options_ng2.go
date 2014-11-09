// +build ignore
package main

// START OMIT

type Dog struct {
	Name string
	Age  int // HL
}

func (d *Dog) SetAge(age int) *Dog { // HL
	d.Age = age // HL
	return d    // HL
}

func NewDog(name string) *Dog {
	d := &Dog{Name: name}
	// Some initialization processes go here.
	return d
}

func main() {
	d := NewDog("Taro").SetAge(7) // HL
}

// END OMIT
