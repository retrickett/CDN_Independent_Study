package main
import("fmt"
	"sync"
	"time"
)

//simulate peer to peer networking
// each node/ peer is a go routine
// one go routine can send a message to a set of other go routines?
// message is a string or an integer

var wg sync.WaitGroup
func stringSliceFromChannel(maxLength int, input <-chan int) []int {
	var results []int
	timeout := time.After(time.Duration(3000) * time.Millisecond)
	time.Sleep(time.Second)
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


func id(channel chan int, id int, msg int, peerprint chan int){
	//goal have case 1 and case 2 wait until case 3 finishes
	defer wg.Done()
	//time.After(time.Duration(3000) * time.Millisecond)
	//fmt.Println("does this print?")
	switch id {
	case 1:
		channel <- msg

		time.Sleep(1*time.Second)
		//time.After(time.Duration(4000) * time.Millisecond)
		//fmt.Println("here?")
		fmt.Println(<-peerprint)
		fmt.Println(1)
		//fmt.Println(stringSliceFromChannel(1,peerprint))

	case 2:
		channel <- msg
		time.Sleep(2*time.Second)
		//fmt.Println("here2")
		//time.After(time.Duration(3000) * time.Millisecond)
		fmt.Println(<-peerprint)
		fmt.Println(2)
		//fmt.Println(stringSliceFromChannel(1,peerprint))
	case 3:
		//fmt.Println("here3")
		msg1 := <-channel
		msg2 := <-channel
		//channel <- msg1 + msg2
		peerprint <- msg1 + msg2
		peerprint <- msg1 + msg2
		//time.After(time.Duration(3000) * time.Millisecond)
		time.Sleep(3*time.Second)
		//wg.Done()
	}
}

func main() {
	//nothing prints when I try to just call slice in the id function so i am trying to add a wait group

	//hint be creative. can have more than 1 channel
	channel := make(chan int,10)
	peerprint := make(chan int, 10)

	wg.Add(3)

	go id(channel,1,5,peerprint)
	//wg.Wait()
	//wg.Add(1)
	go id(channel, 2, 6, peerprint)
	//wg.Wait()
	//wg.Add(1)
	go id(channel, 3, 0, peerprint)
	wg.Wait()
	fmt.Println("sad")



	//wg.Wait()
	//res := stringSliceFromChannel(1,channel)
	//fmt.Println(res)


}
