package main

import (
	"fmt"
)

type errorString struct {
	s string
}

func (e errorString) Error() string {
	//TODO implement me
	return e.s
}

func New(text string) error {
	return &errorString{text}
}

func main() {
	_, e := test()

	fmt.Println(e.errorCode)
	fmt.Println(e.errorMsg)
	e.errorMsg = "yy"
	fmt.Println(e.Error())

}

type commonError struct {
	errorCode int
	errorMsg  string
}

func (e *commonError) Error() string {
	return e.errorMsg
}

func test() (i int, e commonError) {
	e.errorMsg = "xx"
	e.errorCode = 1
	return 0, e
}

type MyError struct {
	err error
	msg string
}

func (e *MyError) Error() string {
	//TODO implement me
	//errors.Unwrap(e)
	//errors.Is(e,e)
	//errors.As(nil, &e)
	return e.err.Error() + e.msg
}
