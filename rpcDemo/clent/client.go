package main

import (
	"log"
	"net/rpc"
)

type Result struct {
	Num int
	Ans float64
}

func main() {
	client, _ := rpc.DialHTTP("tcp", "localhost:1080")
	var result1, result2 Result

	//同步
	if err := client.Call("Cal.Square", 12, &result1); err != nil {
		log.Fatalln("Failed to call Cal.Square. ", err)
	}

	log.Printf("%d^2 = %d\n", result1.Num, result1.Ans)

	//异步
	asyncCall := client.Go("Cal.Square", 11, &result2, nil)
	log.Printf("%d^2 = %d\n", result2.Num, result2.Ans)
	log.Println("------------------------------------")
	log.Println("------------other methods-----------")
	log.Println("------------------------------------")
	<-asyncCall.Done
	log.Printf("%d^2 = %d", result2.Num, result2.Ans)
}
