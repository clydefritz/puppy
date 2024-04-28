package puppy

import (
	"fmt"

	"github.com/clydefritz/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func DogBark() string {
	return dog.WhenGrownUp(Bark())
}

func DogBarks() string {
	return dog.WhenGrownUp(Barks())
}

func From11() {
	fmt.Println("I'm from version 1.1.0")
}

func From12() {
	fmt.Println("I'm from version 1.2.0")
}
