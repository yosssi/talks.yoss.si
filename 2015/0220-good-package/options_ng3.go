// +build ignore
package main

// START OMIT

type Dog struct {
	name    string
	age     int
	sex     string // HL
	species string // HL
}

func (d *Dog) SetSex(sex string) *Dog { // HL
	d.sex = sex // HL
	return d    // HL
}

func (d *Dog) SetSpecies(species string) *Dog { // HL
	d.species = species // HL
	return d            // HL
}

func main() {
	d := NewDog("Taro").SetAge(7).SetSex("male").SetSpecies("Dachshund") // HL
}

// END OMIT
