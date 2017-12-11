package main

import (
	"net"
	"log"
	"io"
	"bufio"
	"fmt"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln()
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		defer conn.Close()

		fmt.Println("cod go here")
		io.WriteString(conn, "i see you")

		conn.Close()
	}

}
