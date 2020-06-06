package interfacefun

import "fmt"

// based on https://npf.io/2014/05/intro-to-go-interfaces/

type Walker interface {
	Walk(miles int)
	Distance() int
	WalkerName() string
}

type Dog struct {
	Name           string
	DistanceWalked int
}

type Cat struct {
	Name           string
	DistanceWalked int
}

func (d Dog) Walk(miles int) {
	// here d is a COPY of the sam object
	d.DistanceWalked += miles // therefore this line doesn't change the original sam's distance
	fmt.Printf("%s (%p) is walking %v miles, %v overall walked\n", d.Name, &d, miles, d.DistanceWalked)
}

func (d Dog) WalkerName() string {
	return d.Name
}

func (d Dog) Distance() int {
	return d.DistanceWalked
}

func (c *Cat) Walk(miles int) {
	// here c is the pointer to the original tim object
	c.DistanceWalked += miles // therefore this line changes the original tim's distance
	fmt.Printf("%s (%p) is walking %v miles, %v overall walked\n", c.Name, c, miles, c.DistanceWalked)
}

// implemented as value receiver because this function does
// not change anything in cat
func (c Cat) Distance() int {
	return c.DistanceWalked
}

func (c Cat) WalkerName() string {
	return c.Name
}

/*
what "actually" happens when assigning a struct to an interface:

data := c
w := Walker{
    type: &InterfaceType{
              valtype: &typeof(c),
              func0: &Dog.Walk
          }
    data: &data
}
*/
func LongWalk(walker Walker) {
	walker.Walk(500)
	fmt.Printf("%s's distance after 1st walk %v\n",
		walker.WalkerName(),
		walker.Distance())

	copyOfWalker := walker
	copyOfWalker.Walk(25)

	fmt.Printf("%s's distance after 2nd walk %v\n",
		walker.WalkerName(),
		walker.Distance())
}
