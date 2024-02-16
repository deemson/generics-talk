package examples

import (
	"fmt"
	"testing"
)

type Animal interface {
	MakeSound() string
}

type Cat struct{}

func (c Cat) MakeSound() string {
	return "meow"
}

type Dog struct{}

func (d Dog) MakeSound() string {
	return "woof"
}

func makeAllAnimalsDoSounds(animals []Animal) {
	for _, animal := range animals {
		fmt.Println(animal.MakeSound())
	}
}

func TestListOfAnimals(t *testing.T) {
	//cats := []Cat{{}, {}}
	//makeAllAnimalsDoSounds(cats)
}

type Box[T any] struct {
	Value T
}

func makeAnimalInTheBoxDoSound(box Box[Animal]) {
	fmt.Println(box.Value.MakeSound())
}

func TestAnimalInABox(t *testing.T) {
	//catInTheBox := Box[Cat]{Cat{}}
	//makeAnimalInTheBoxDoSound(catInTheBox)
}
