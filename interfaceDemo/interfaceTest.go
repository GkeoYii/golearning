package interfaceDemo

import "fmt"

type Person interface {
	position() string
}

type Teacher struct {
	work string
}

type Cook struct {
	work string
}

func (t Teacher) position() string {
	return "teacher"
}

func (c Cook) position() string {
	return "cook"
}

func Println(p Person) {
	fmt.Println("it is ", p.position())
}
