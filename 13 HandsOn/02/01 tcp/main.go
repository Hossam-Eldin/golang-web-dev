package main

import (
	"net"
	"log"
	"io"
)

func main() {
	li, err := net.Listen("tcp",":8080")
	if err != nil{
		log.Fatalln()
	}
	defer li.Close()
	for  {
	 conn,err:=li.Accept()
		if err != nil{
			log.Fatalln(err)
			continue
		}

		io.WriteString(conn,"i see you")

		conn.Close()
	}
}
