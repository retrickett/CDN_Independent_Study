package main
import("fmt"
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


func send(peer1 chan int, msg int) {
	peer1 <- msg
}

func receive(peer1 chan int, peer2 chan int) {
	msg := <-peer1
	peer2 <- msg
}

func main() {
	peer1 := make(chan int)
	peer2 := make(chan int)
	go send(peer1,5)
	go receive(peer1, peer2)
	go send(peer2, 10)
	go receive(peer2, peer1)

	peer1res := stringSliceFromChannel(1,peer1)
	peer2res := stringSliceFromChannel(1,peer2)
	fmt.Println(peer1res)
	fmt.Println(peer2res)


}
