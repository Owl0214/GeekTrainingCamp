package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	queue1 := make(chan string, 10)
	stopCh := make(chan interface{}, 1)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer func() {
			cancel()
			wg.Done()
		}()
		producer(queue1, ctx, "producer1")
	}()

	wg.Add(1)
	go func() {
		defer func() {
			cancel()
			wg.Done()
		}()
		producer(queue1, ctx, "producer2")
	}()

	wg.Add(1)
	go func() {
		defer func() {
			cancel()
			wg.Done()
		}()
		consumer(queue1, stopCh)
	}()

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		for {
			select {
			case <-ctx.Done():
				stopCh <- true
				return
			case <-time.After(3 * time.Second):
				fmt.Println("queue has product:", len(queue1))
				if len(queue1) == cap(queue1) {
					cancel()
				}
			}
		}
	}()
	wg.Wait()
}

func producer(queue chan string, ctx context.Context, name string) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			// 队列满了，不生产了
			return
		case <-time.After(time.Second):
			i++
			fmt.Printf("Producer %v produced:%v\n", name, i)
			queue <- name + "'s production" + fmt.Sprint(i)
			if len(queue) == 10 {
				ctx.Value("full")
			}
		}
	}
}

func consumer(queue chan string, stopCh chan interface{}) {
	i := 0
	for {
		select {
		case <-stopCh:
			j := 0
			for len(queue) > 0 {
				time.After(time.Second)
				result := <-queue
				j++
				fmt.Printf("Consumer get production %v times,and result is %v finally and will return\n", j, result)
			}
			return
		case <-time.After(time.Second):
			result := <-queue
			i++
			fmt.Printf("Consumer try to get production %v times,and result is %v\n", i, result)
		}
	}
}
