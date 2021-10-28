package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

var channel1 = make(chan int, 10)
var channel2 = make(chan int, 10)
var channel3 = make(chan int, 10)
var channel4 = make(chan int, 10)
var channel5 = make(chan int, 10)
var channel6 = make(chan int, 10)

func containsInt (s []int, y int) bool{
	var x = false
	for i:=0;i< len(s);i++{
		if s[i] == y{
			x= true
		}
	}
	return x
}

func round(channel chan int, numreached []int){
	for len(numreached) <6 {
		rand.Seed(time.Now().UnixNano())
		peer := (rand.Intn(7-1)+1)
		switch peer {
		case 2:
			if containsInt(numreached, 2) == false {
				numreached = append(numreached, 2)
				data := <-channel
				channel <- data
				channel2 <- data

			}
			gossipSpread(numreached, 2, channel2)

		case 3:
			if containsInt(numreached, 3) == false {
				numreached = append(numreached, 3)
				data := <-channel
				channel <- data
				channel3 <- data

			}
			gossipSpread(numreached, 3, channel3)
		case 4:
			if containsInt(numreached, 4) == false {
				numreached = append(numreached, 4)
				data := <-channel
				channel <- data
				channel4 <- data

			}
			gossipSpread(numreached, 4, channel4)
		case 5:
			if containsInt(numreached, 5) == false {
				numreached = append(numreached, 5)
				data := <-channel
				channel <- data
				channel5 <- data

			}
			gossipSpread(numreached, 5, channel5)

		case 6:
			if containsInt(numreached, 6) == false {
				numreached = append(numreached, 6)
				data := <-channel
				channel <- data
				channel6 <- data

			}
			gossipSpread(numreached, 6, channel6)

		}
		//fmt.Println(peer)
	}
	fmt.Println(numreached)
}

func gossipSpread(numreached []int, id int, channel chan int){
	defer wg.Done()
	switch id{
	case 1 :
		round(channel1, numreached)
	case 2:
		round(channel2, numreached)
	case 3:
		round(channel3, numreached)
	case 4:
		round(channel4, numreached)
	case 5:
		round(channel5, numreached)
	case 6:
		round(channel6, numreached)
		}

}
func main(){

	var numreached []int

	numreached = append(numreached,1)
	channel1 <-5
	wg.Add(6)
	go gossipSpread(numreached,1, channel1)
	go gossipSpread(numreached,2,channel2)
	go gossipSpread(numreached,3,channel3)
	go gossipSpread(numreached,4,channel4)
	go gossipSpread(numreached,5,channel5)
	go gossipSpread(numreached,6,channel6)
	wg.Wait()


	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
	fmt.Println(<-channel3)
	fmt.Println(<-channel4)
	fmt.Println(<-channel5)
	fmt.Println(<-channel6)


}
