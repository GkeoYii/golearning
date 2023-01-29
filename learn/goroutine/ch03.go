package main

import (
	"fmt"
	"sync"
	"time"
)

// sync.Mutex     .Lock() .Unlock()
// sync.RWMutex    .RLock()   .RUnLock()
// sync.WaitGroup   .add()  .Done()   .Wait()
// sync.Once
// sync.Cond  .Wait()   .Signal()  .broadCast()
var (
	sum     int
	mutex   sync.Mutex //互斥锁
	rwMutex sync.RWMutex
)

func add(i int) {
	mutex.Lock()
	defer mutex.Unlock()
	sum += i
}

func readSum() int {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	b := sum
	return b
}

func run() {
	var wg sync.WaitGroup
	wg.Add(110)
	for i := 0; i < 100; i++ {
		//检查是否存在资源竞争，go build/run/test -race
		//go add(10)
		go func() {
			defer wg.Done()
			add(10)
		}()
	}
	for i := 0; i < 10; i++ {
		//go fmt.Println("sum = ", readSum())
		go func() {
			defer wg.Done()
			fmt.Println("sum = ", readSum())
		}()
	}
	//time.Sleep(2 * time.Second)
	//等待计数器值为0结束
	wg.Wait()
}

func numberSum(intChan chan int, sumChan chan int) {
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		sum := 0
		for i := 1; i <= num; i++ {
			sum += i
		}
		fmt.Println("sum : =", sum)
		sumChan <- sum
	}

}
func putNum(intChan chan int) {
	for i := 1; i <= 8; i++ {
		intChan <- i
	}
	//关闭
	close(intChan)
}

func main() {

	//让主协程休眠，执行从协程
	//go fmt.Println("this is goroutine")
	//fmt.Println("this is main func")
	//time.Sleep(time.Second)

	//channelTest2()
	//run()
	//race()

	num := 8
	wg := sync.WaitGroup{}
	wg.Add(8)

	numChan := make(chan int, num)
	sumChan := make(chan int, num)
	go putNum(numChan)
	for i := 0; i < 8; i++ {
		go func() {
			defer wg.Done()
			numberSum(numChan, sumChan)
		}()
	}

	go func() {
		defer wg.Done()
		fmt.Println()
		for {
			sum, ok := <-sumChan
			if !ok {
				break
			}
			fmt.Println("main sum", sum)
		}
	}()

	wg.Wait()
}

func channelTest() {

	ch := make(chan string) //无缓存channel
	//cache := make(chan int, 5)      //有缓存channel

	//onlySend := make(chan<- int)    //只能发送
	//onlySend <- 6

	//onlyReceive := make(<-chan int) //只能接收
	//i := <-onlyReceive
	go func() {
		fmt.Println("this is slave1")
		ch <- "slave1 complete"
	}()
	fmt.Println("this is main")
	//阻塞主协程执行，只有协程执行后，才执行
	v := <-ch
	fmt.Println("accept v", v)
}

// 多路复用
func channelTest2() {
	firstCh := make(chan string)
	secondCh := make(chan string)
	thirdCh := make(chan string)

	go func() {
		firstCh <- downloadFile("firstCh")
	}()
	go func() {
		secondCh <- downloadFile("secondCh")
	}()
	go func() {
		thirdCh <- downloadFile("thirdCh")
	}()

	select {
	case filePath := <-firstCh:
		fmt.Println(filePath)
	case filePath := <-secondCh:
		fmt.Println(filePath)
	case filePath := <-thirdCh:
		fmt.Println(filePath)
	}
}

func downloadFile(chanName string) string {
	time.Sleep(time.Second)
	return chanName + ":filePath"
}

func race() {
	cond := sync.NewCond(&sync.Mutex{})
	//var m sync.Map
	var wg sync.WaitGroup
	wg.Add(11)
	for i := 0; i < 10; i++ {
		go func(num int) {
			defer wg.Done()
			fmt.Println(num, "号已经就位")
			cond.L.Lock() // 互斥锁
			cond.Wait()   // 阻塞当前协程
			fmt.Println(num, "号开始跑...")
			cond.L.Unlock()
		}(i)
	}

	time.Sleep(2 * time.Second)

	go func() {
		defer wg.Done()
		fmt.Println("裁判已经就位，准备发令枪")
		fmt.Println("比赛开始，大家准备跑")
		//cond.Signal() //唤醒一个等待时间最长的协程
		//time.Sleep(2 * time.Second)
		cond.Broadcast() //唤醒全部协程
	}()
	wg.Wait()
}
