package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	username := fetchUser()
	respCh := make(chan any, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go fetchUserLikes(username, respCh, wg)
	go fetchUserMatch(username, respCh, wg)
	wg.Wait() // block until 2 wg.Done()
	close(respCh)

	for resp := range respCh {
		fmt.Println("resp: ", resp)
	}

	//fmt.Println("likes: ", likes)
	//fmt.Println("match: ", match)
	fmt.Println("took: ", time.Since(start))
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)
	return "agg"
}

func fetchUserLikes(username string, respCh chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)
	respCh <- 11
	wg.Done()
}

func fetchUserMatch(username string, respCh chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	respCh <- "ANNA"
	wg.Done()
}
