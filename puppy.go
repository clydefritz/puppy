package puppy

import "github.com/clydefritz/dog"

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
