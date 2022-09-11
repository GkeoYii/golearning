package main

import (
	"log"
	"math"
	"net/http"
	"net/rpc"
)

type Result struct {
	Num int
	Ans float64
}
type Cal int

// Square ：func (t *T) MethodName(argType T1, replyType *T2) error
func (cal *Cal) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = math.Pow(float64(num), 2)
	return nil
}

// Remote Procedure Call
func main() {
	rpc.Register(new(Cal))
	rpc.HandleHTTP()
	log.Printf("Serving RPC server on port %d", 1080)
	if err := http.ListenAndServe(":1080", nil); err != nil {
		log.Fatalln("Error serving", err)
	}
}
