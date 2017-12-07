package main

import (
	"net"
	"log"
	"io/ioutil"
	"fmt"
)

func main() {

	conn, err := net.Dial("tcp",":8080")
	if err != nil{
		log.Fatalln(err)
	}

	defer conn.Close()

	bs,err := ioutil.ReadFile(conn)
	if err !=nil{
		log.Fatalln(err)
	}

	fmt.Println(string(bs))

}
