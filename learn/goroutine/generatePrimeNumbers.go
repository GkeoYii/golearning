package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	xMutex sync.Mutex
)

func main() {

	num := 1000000
	var nums []int
	var wg sync.WaitGroup
	wg.Add(num)
	fmt.Println("start time :", time.Now())

	for i := 1; i <= num; i++ {
		go func(x int) {
			defer wg.Done()
			y, ok := isPrimeNumber(x)

			if ok {
				xMutex.Lock()
				defer xMutex.Unlock()
				nums = append(nums, y)
				//fmt.Println("y = ", y)
			}
		}(i)

		//start time : 2023-01-29 16:56:03.912076 +0800 CST m=+0.000130835
		//end time : 2023-01-29 16:56:18.848977 +0800 CST m=+14.937087585
		//y, ok := generateNatural(i)
		//if ok {
		//	nums = append(nums, y)
		//}

	}
	wg.Wait()
	fmt.Println("end time :", time.Now())
}

func isPrimeNumber(num int) (int, bool) {
	if num == 0 || num == 1 {
		return 0, false
	}
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return 0, false
		}
	}
	return num, true
}
