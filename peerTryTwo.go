package main
import("fmt"
	"sync"
	"time"
)

//simulate peer to peer networking
// each node/ peer is a go routine
// one go routine can send a message to a set of other go routines?
// message is a string or an integer

func stringSliceFromChannel(maxLength int, input <-chan int) []int {
	var results []int
	timeout := time.After(time.Duration(80) * time.Millisecond)

	for {
		select {
		case int := <-input:
			results = append(results, int)

			if len(results) == maxLength {
				//fmt.Println("Got all results")
				return results
			}
		case <-timeout:
			fmt.Println("Timeout!")
			return results
		}
	}
}


//func send(peer1 chan int, msg int) {
//	peer1 <- msg
//}

func receive(peer1 chan int, peer2 chan int) {
	msg := <-peer1
	peer2 <- msg
}

func send(peer chan int, msg int){
	peer <-msg
	//time.Sleep(10)
	fmt.Println(<-peer)
}

func peer3(peer chan int){
	msg1 := <-peer
	msg2 := <-peer
	peer <-msg1 + msg2
	fmt.Println(<- peer)
	fmt.Println("WHATTTTTUPPPP")
}

func id(channel chan int, id int, msg int, wg *sync.WaitGroup){
	//defer wg.Done()
	time.Sleep(time.Second)
	switch id {
	case 1:
		channel <- msg
	case 2:
		channel <- msg
	case 3:
		msg1 := <-channel
		msg2 := <-channel
		channel <- msg1 + msg2
	}
}

func main() {
	//peer1 := make(chan int)
	//peer2 := make(chan int)
	var wg sync.WaitGroup

	channel := make(chan int)
	//go send(channel,5)
	//go send(channel, 10)
	//go peer3(channel)
	//fmt.Println("YO")
	//wg.Add(1)
	go id(channel, 1, 5,&wg)
	//wg.Add(1)
	go id(channel, 2, 10,&wg)
	//wg.Add(1)
	go id(channel, 3, 0,&wg)
	wg.Wait()
	res := stringSliceFromChannel(3,channel)
	fmt.Println(res)
	//fmt.Println(<-channel)

	//go send(peer1,5)
	//go receive(peer1, peer2)
	//go send(peer2, 10)
	//go receive(peer2, peer1)

	//peer1res := stringSliceFromChannel(1,peer1)
	//peer2res := stringSliceFromChannel(1,peer2)
	//fmt.Println(peer1res)
	//fmt.Println(peer2res)


}
