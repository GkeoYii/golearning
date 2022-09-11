package main

import (
	"log"
	"math"
)

type Result struct {
	Num int
	Ans float64
}
type Cal int

//first
//func (cal Cal) Square(num int) *Result {
//	return &Result{
//		Num: num,
//		Ans: math.Pow(float64(num), 2),
//	}
//}

// Square ：func (t *T) MethodName(argType T1, replyType *T2) error
func (cal *Cal) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = math.Pow(float64(num), 2)
	return nil
}

func main() {
	//first
	//cal := new(Cal)
	//result := cal.Square(12)
	//log.Println("%d^2 = %d", result.Num, result.Ans)

	cal := new(Cal)
	var result Result
	cal.Square(11, &result)
	log.Println("%d^2 = %d", result.Num, result.Ans)
}
