package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type message struct{
	Name string
}
type receive struct{
	Name string
}

func main(){
	//for true {
	//	arg_len := len(os.Args[1:])
	//	for i:=0; i<arg_len; i++{
	//		fmt.Println(os.Args[i])
	//	}
	//}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a message:")
	text, _ := reader.ReadString('\n')
	//fmt.Print(text)
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)

	err := enc.Encode(message{text})
	if err != nil{
		log.Fatal("uh oh",err)
	}

	var get receive
	err = dec.Decode(&get)
	if err!= nil{
		log.Fatal("uh oh 2", err)
	}
	fmt.Println(get.Name)
}
