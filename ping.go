package main

import (
	"fmt"
	"os/exec"
	"time"
)

func pingServer(site string, c chan string){
	prg := "ping"
	cflag := "-c"
	npings := "1"
	cmd := exec.Command(prg, cflag, npings, site)
	stdout, err := cmd.Output()

	if err != nil{
		c <- err.Error()
	}

	c <- string(stdout)
}

//sample slicing code from website 
func stringSliceFromChannel(maxLength int, input <-chan string) []string {
	var results []string
	timeout := time.After(time.Duration(80) * time.Millisecond)

	for {
		select {
		case str := <-input:
			results = append(results, str)

			if len(results) == maxLength {
				fmt.Println("Got all results")
				return results
			}
		case <-timeout:
			fmt.Println("Timeout!")
			return results
		}
	}
}

func main(){
	messages := make(chan string)
	website := "www.amazon.com"
	for i:=0;i<4;i++{
		go pingServer(website, messages)
		//fmt.Println(<-messages)
	}
	results := stringSliceFromChannel(3,messages)
	fmt.Println("Results: =%v\n",results)
	fmt.Println("hi")
	//fmt.Println(<-messages)
	//pingServer(website, messages)
	//fmt.Println(<-messages)
}
