package main

import (
	"fmt"
	"sync"
	_ "time"
	"os/exec"
)

func pings(webAddress string, wg *sync.WaitGroup) <- chan string{
	defer wg.Done()
	output := make(chan string)
	cmd, _ := exec.Command("ping", "-c1", webAddress).Output()
	//output <- string(cmd)
	fmt.Println(string(cmd))
	//close(output)
	return output
}

func main() {
	var websites = []string{"google.com", "amazon.com", "yahoo.com"}
	var wg sync.WaitGroup
	for _, website := range websites{
		wg.Add(1)
		go pings(website, &wg)
	}
	wg.Wait()
}


