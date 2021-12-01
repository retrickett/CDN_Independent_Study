package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
)

type message struct{
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
	//tcp
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Connection error", err)
	}

	//var network bytes.Buffer
	enc := gob.NewEncoder(conn)


	err = enc.Encode(message{text})
	if err != nil{
		log.Fatal("uh oh",err)
	}

	//conn.Close()
	//fmt.Println(get.Name)
}
