// +build ignore
package main

// START OMIT

type Dog struct {
	Name    string
	Age     int
	Sex     string  // HL
	Species string  // HL
	Length  float32 // HL
	Weight  float32 // HL
}

func main() {
	d := NewDog("Taro").SetAge(7).SetSex("male").SetSpecies("Dachshund").SetLength(8.5).SetWeight(20.5) // HL
}

// END OMIT
