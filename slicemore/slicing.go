package main

import "fmt"

func main() {
	a := []byte{'g', 'o', 'l', 'a', 'n', 'g'}

	b := a[:2]

	b[0] = 1 // this supposes to take effect in both a and b because of the same underlying array

	fmt.Printf("before append a=%v, b=%v\n", a, b) // both a and b have 1 as first element

	a = append(a, 'x') // accidentally append beyond one slice capacity. imagine this happens when tons of slices working in a loop.

	b[1] = 2
	fmt.Printf("after append a=%v, b=%v\n", a, b) // now a and b are silently pointing to totally independent underlying arrays

}
