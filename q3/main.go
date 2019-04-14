package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func a1(wg *sync.WaitGroup) {
	q := rand.Intn(3)
	sleepTime := 1 + q
	time.Sleep(time.Duration(sleepTime) * time.Second)
	fmt.Println("a1 finished!")
	wg.Done()
}

func a2(wg *sync.WaitGroup) {
	q := rand.Intn(5)
	sleepTime := 2 + q
	time.Sleep(time.Duration(sleepTime) * time.Second)
	fmt.Println("a2 finished!")
	wg.Done()
}

func a3(wg *sync.WaitGroup) {
	q := rand.Intn(2)
	sleepTime := 3 + q
	time.Sleep(time.Duration(sleepTime) * time.Second)
	fmt.Println("a3 finished!")
	wg.Done()
}

func b1(wg *sync.WaitGroup) {
	q := rand.Intn(2)
	sleepTime := 5 + q
	time.Sleep(time.Duration(sleepTime) * time.Second)
	fmt.Println("b1 finished!")
	wg.Done()
}

func b2(wg *sync.WaitGroup) {
	q := rand.Intn(3)
	sleepTime := 8 + q
	time.Sleep(time.Duration(sleepTime) * time.Second)
	fmt.Println("b2 finished!")
	wg.Done()
}

func b3(wg *sync.WaitGroup) {
	q := rand.Intn(2)
	sleepTime := 4 + q
	time.Sleep(time.Duration(sleepTime) * time.Second)
	fmt.Println("b3 finished!")
	wg.Done()
}

func c1(wg *sync.WaitGroup) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("c1 finished!")
	wg.Done()
}

func c2(wg *sync.WaitGroup) {
	q := rand.Intn(4)
	sleepTime := (1.5 + float64(q)*0.5) * 1000
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	fmt.Println("c2 finished!")
	wg.Done()
}

func c3(wg *sync.WaitGroup) {
	q := rand.Intn(3)
	sleepTime := 6 + q
	time.Sleep(time.Duration(sleepTime) * time.Second)
	fmt.Println("c3 finished!")
	wg.Done()
}

func d() {
	fmt.Println("d finished!")
}

func main() {
	var wg1, wg2, wg3, wgd sync.WaitGroup
	{
		wg1.Add(1)
		go a1(&wg1)
	}
	{
		wg2.Add(1)
		go a2(&wg2)
	}
	{
		wg3.Add(1)
		go a3(&wg3)
	}
	{
		wg1.Add(1)
		go b1(&wg1)
	}
	{
		wg2.Add(1)
		go b2(&wg2)
	}
	{
		wg3.Add(1)
		go b3(&wg3)
	}
	{
		wgd.Add(3)
		wg1.Wait()
		go c1(&wgd)
		wg2.Wait()
		go c2(&wgd)
		wg3.Wait()
		go c3(&wgd)
	}
	{
		wgd.Wait()
		go d()
	}
	time.Sleep(15 * time.Second)
	fmt.Println(time.Now().Hour())
}
