package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//for {
	//	select {
	//	case <-done:
	//		return
	//	default:
	//		fmt.Println("TODO")
	//	}
	//}

	//for _, s := range []int{} {
	//	select {
	//	case <-done:
	//		return
	//	case resultCh <- s:
	//		fmt.Println("TODO")
	//	}
	//}
	//timeOutTest()

	//coms := buy(10) //采购10套配件
	//
	//phones := build(coms) //组装10部手机
	//
	//packs := pack(phones) //打包它们以便售卖
	//
	////输出测试，看看效果
	//

	coms := buy(100)

	phone1 := build(coms)
	phone2 := build(coms)
	phone3 := build(coms)

	phones := merge(phone1, phone2, phone3)
	packs := pack(phones)
	for p := range packs {
		fmt.Println(p)
	}
}

func timeOutTest() {
	result := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		result <- "data"
	}()

	select {
	case v := <-result:
		fmt.Println(" result = ", v)
	case <-time.After(3 * time.Second):
		fmt.Println("request timeout")
	}
}

func buy(n int) <-chan string {

	out := make(chan string)

	go func() {

		defer close(out)

		for i := 1; i <= n; i++ {

			out <- fmt.Sprint("配件", i)

		}

	}()

	return out

}
func build(in <-chan string) <-chan string {

	out := make(chan string)

	go func() {

		defer close(out)

		for c := range in {

			out <- "组装(" + c + ")"

		}

	}()

	return out

}

func pack(in <-chan string) <-chan string {

	out := make(chan string)

	go func() {

		defer close(out)

		for c := range in {

			out <- "打包(" + c + ")"

		}

	}()

	return out

}

func merge(ins ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	p := func(in <-chan string) {
		defer wg.Done()
		for c := range in {
			out <- c
		}
	}

	wg.Add(len(ins))

	for _, cs := range ins {
		go p(cs)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
