package main
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var channel1in = make(chan int, 10)
var channel2in = make(chan int, 10)
var channel3in = make(chan int, 10)
var channel4in = make(chan int, 10)
var channel5in = make(chan int, 10)
var channel6in = make(chan int, 10)
//var channel1out = make(chan int, 10)
////var channel2out = make(chan int, 10)
//var channel3out = make(chan int, 10)
//var channel4out = make(chan int, 10)
//var channel5out = make(chan int, 10)
//var channel6out = make(chan int, 10)


//make a function that just handles the spreading
//make a round function that is called saying if you are in numreached call gossip2

func containsInt2 (s []int, y int) bool{
	var x = false
	for i:=0;i< len(s);i++{
		if s[i] == y{
			x= true
		}
	}
	return x
}
func returnChan(id int) chan int{
	switch id{
	case 1:
		return channel1in

	case 2:
		return channel2in
	case 3:
		return channel3in

	case 4:
		return channel4in

	case 5:
		return channel5in

	default:
		return channel6in
	}
}

func rounds2(numreached []int, id int){
	if containsInt2(numreached, id){
		gossipSpread2(numreached, id, returnChan(id))
	}
	if len(numreached)==6{
		println("done")
	}
}



func gossipSpread2(numreached []int, id int, from chan int){
	//randomly pick a routine. if the numreached2 just has 1 in it, then put channel1 in to be whats giving it
	rand.Seed(time.Now().UnixNano())
	routine := (rand.Intn(7-1)+1)
	// we want all peers to be in each round so dont call gossipspread in the cases
	if containsInt2(numreached, routine)==false{
		numreached = append(numreached, routine)
		data:= <-from
		from <-data
		returnChan(routine) <- data
	}
	rounds2(numreached, id)

	//switch routine{
	//case 1:
	//case 2:
	//case 3:
	//case 4:
	//case 5:
	//case 6:
	//	}
}

func main(){


	// we want 2 channels between each go routine

	var numreached2 []int

	numreached2 = append(numreached2,1)
	channel1in <-5
	wg.Add(6)
	go rounds2(numreached2,1)
	go rounds2(numreached2,2)
	go rounds2(numreached2,3)
	go rounds2(numreached2,4)
	go rounds2(numreached2,5)
	go rounds2(numreached2,6)

	//go gossipSpread2(numreached2,1, channel1in)
	//go gossipSpread2(numreached2,2, channel1in)
	//go gossipSpread2(numreached2,3, channel1in)
	//go gossipSpread2(numreached2,4, channel1in)
	//go gossipSpread2(numreached2,5, channel1in)
	//go gossipSpread2(numreached2,6, channel1in)
	wg.Wait()


	fmt.Println(<-channel1in)
	fmt.Println(<-channel2in)
	fmt.Println(<-channel3in)
	fmt.Println(<-channel4in)
	fmt.Println(<-channel5in)
	fmt.Println(<-channel6in)


}
