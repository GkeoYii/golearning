package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	//stopCh := make(chan bool) //停止watchdog
	//go func() {
	//	defer wg.Done()
	//	watchDogByChannel(stopCh, "[watch 🐶]")
	//}()
	//time.Sleep(5 * time.Second)
	//stopCh <- true //发送指令

	//context   .WithCancel()  WithDeadline() WithTimeout() WithValue(parent Context, key, val interface{})

	//ctx, stop := context.WithCancel(context.Background())
	//go func() {
	//	defer wg.Done()
	//	watchDogByContext(ctx, "[watch 🐶]")
	//}()
	//time.Sleep(5 * time.Second)
	//stop() //执行取消方法

	//ctx, stop := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	ctx, stop := context.WithTimeout(context.Background(), 5*time.Second)
	defer stop()

	go func() {
		defer wg.Done()
		watchDogByContext(ctx, "[watch 🐶 no.1]")
	}()

	go func() {
		defer wg.Done()
		watchDogByContext(ctx, "[watch 🐶 no.2]")
	}()

	wg.Wait()
}

func watchDogByChannel(stopCh chan bool, name string) {
	for {
		select {
		case <-stopCh: //判断是否写入停止指令
			fmt.Println(name, "停止指令已收到，马上停止...")
			return //结束
		default:
			fmt.Println(name, "正在监控...")
		}
		time.Sleep(1 * time.Second)
	}

}

func watchDogByContext(ctx context.Context, name string) {

	for {
		select {
		case <-ctx.Done(): //判断是否调用取消方法
			fmt.Println(name, "停止指令已收到，马上停止...")
			return //结束
		default:
			fmt.Println(name, "正在监控...")
		}
		time.Sleep(1 * time.Second)
	}
}
