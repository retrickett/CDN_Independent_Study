package main

import (
	"fmt"
	"sync"
	"time"
)

func send(pipeline chan int, value int, wg *sync.WaitGroup) {
	//defer wg.Done()
	pipeline <- value
}
//this function gets the content from channel and creates and array
func intSliceFromChannel(maxLength int, input <-chan int) []int {
	var results []int
	timeout := time.After(time.Duration(80) * time.Millisecond)
	for {
		select {
		case int := <-input:
			results = append(results, int)
			if len(results) == maxLength {
				return results
			}
		case <-timeout:
			fmt.Println("Timeout!")
			return results
		}
	}
}

//takes array that was created above and sums the integers in it
func sum(result []int, pipeline chan int, wg *sync.WaitGroup) {
	//defer wg.Done()
	fmt.Println("executing")
	sum := 0
	for _, v := range result{
		sum += v
	}
	fmt.Println("Sum = ", sum)
}

func main() {
	peer1 := 1
	peer2 := 2
	pipeline := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1) //adds 1 to wg counter
	//go send(pipeline, peer1, &wg)
	go func() {
		send(pipeline, peer1, &wg)
		wg.Done() //decrements the wg counter
	}()
	wg.Add(1)
	//go send(pipeline, peer2, &wg)
	go func(){
		send(pipeline, peer2, &wg)
		wg.Done()
	}()
	results := intSliceFromChannel(2,pipeline)
	fmt.Println(results)
	wg.Add(1)
	//go sum(results, pipeline, &wg)
	go func() {
		sum(results, pipeline, &wg)
		wg.Done()
	}()
	wg.Wait()
}

//to solve issues, I worked on adding print statements in to see what code was executing and what code wasn't
// if I knew that something was executing, I would print everything up until that point
// with the waitgroup, wg.Add(x) I was just moving things around to see what would execute and what would not