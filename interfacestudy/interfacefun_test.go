package interfacefun

import (
	"fmt"
	"testing"
)

func TestObjectReceiver(t *testing.T) {
	samTheDog := Dog{Name: "sam"}

	// sam is an object. the following call is actually translate implicitly to
	// (&sam).Walk(10)
	fmt.Printf("sam's memory address %p\n", &samTheDog)
	samTheDog.Walk(10)

	// manuall add the 10 miles walked
	samTheDog.DistanceWalked += 10

	LongWalk(samTheDog)

	fmt.Printf("sam's overall distance walked %v\n", samTheDog.DistanceWalked)
}

func TestPointerReceiver(t *testing.T) {
	timTheCat := &Cat{Name: "tim"}

	fmt.Printf("tim's memory address %p\n", timTheCat)
	timTheCat.Walk(10)

	LongWalk(timTheCat)

	fmt.Printf("tim's overall distance walked %v\n", timTheCat.DistanceWalked)
}
