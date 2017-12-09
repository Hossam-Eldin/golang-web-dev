package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	li, err := net.Listen("tcp",":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer  li.Close()

	for {
		conn , err := li.Accept()
		if err !=nil{
			log.Fatalln(err.Error())
			continue
		}

		go handel(conn)
	}
	
}

func handel( conn  net.Conn)  {
	defer conn.Close()

	//01 request
	request(conn)
	//respond
	respond(conn)
}

func request(conn net.Conn)  {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		ln :=scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			m :=strings.Fields(ln)[0]
			u := strings.Fields(ln)[1] // uri
			fmt.Println("***METHOD", m)
			fmt.Println("***URI", u)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func respond(conn net.Conn)  {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}