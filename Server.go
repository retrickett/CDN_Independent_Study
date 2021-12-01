package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)
type message2 struct{
	Name string
}

func connectionHandler(connection net.Conn){
	dec := gob.NewDecoder(connection)
	var in message2
	err := dec.Decode(&in)
	if err!=nil{
		log.Fatal("error decoding", err)
	}
	fmt.Println("Message from Client:")
	fmt.Println(in.Name)

	connection.Close()
}

func main(){
	fmt.Println("Starting server")
	ln, err := net.Listen("tcp", ":8080")
	if err!=nil{
		log.Fatal("error connecting to tcp",err)
	}

	for {
		connection, err := ln.Accept()
		if err!=nil{
			log.Fatal("error accepting connection",err)
		}
		go connectionHandler(connection)
	}


}