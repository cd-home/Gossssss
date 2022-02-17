package main

import (
	"fmt"
	"testing"
)

// Animal Go 遵循组合，而不是子类或者继承
// Base type
type Animal struct {
	Name     string
	IsMammal bool
}

// Speak 行为 common behavior
func (a *Animal) Speak() {
	fmt.Println("My Name is", a.Name, ", it is", a.IsMammal, ", I am a mammal")
}


type Dog struct {
	// embedding
	Animal
	PackFactor int
}

func (d *Dog) Speak() {
	fmt.Println("My Name is", d.Name, ", it is", d.IsMammal, ", I am a mammal", d.PackFactor)
}


type Cat struct {
	Animal
	ClimbFactor int
}

func (c *Cat) Speak() {
	fmt.Println("My Name is", c.Name, ", it is", c.IsMammal, ", I am a mammal", c.ClimbFactor)
}


func TestMain(test *testing.T) {
	animal := Animal{
		Name:     "Mino",
		IsMammal: true,
	}
	animal.Speak()

	d := Dog{
		Animal: Animal{
			Name: "JO",
			IsMammal: true,
		},
		PackFactor: 12,
	}
	d.Speak()

	c := Cat{
		Animal: Animal{
			Name: "JO",
			IsMammal: true,
		},
		ClimbFactor: 10,
	}
	c.Speak()
}

func TestEmbeddingNotinheritance(t *testing.T) {
	// Can not compile Dog is Dog Cat is Cat
	/*
	animals := []Animal{ 
			Dog{
				Animal: Animal{
					 Name: "Fido",
					  IsMammal: true,
				}, PackFactor: 5,
			},
			Cat{
				Animal: Animal{ 
					Name: "Milo",
					IsMammal: true,
				},
			ClimbFactor: 4, },
	}
	for _, animal := range animals {
		 animal.Speak()
	}
	*/
	t.Log("Embedding isn’t the same as inheritance and this is the pattern I need to stay away from. A Dog is a Dog, a Cat a Cat, and an Animal an Animal")
}

// define the common method set of behavior 
type Speaker interface {
	Speak()
}

func TestInterface(t *testing.T) {
	// Can
	speakers := []Speaker{ 
			&Dog{
				Animal: Animal{
					 Name: "Fido",
					  IsMammal: true,
				}, PackFactor: 5,
			},
			&Cat{
				Animal: Animal{ 
					Name: "Milo",
					IsMammal: true,
				},
			ClimbFactor: 4, },
	}
	for _, speaker := range speakers {
		 speaker.Speak()
	}
}