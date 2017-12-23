package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"io"
	"fmt"
)

func main() {
	c:=getCode("test@example.com")
	fmt.Println("ORGINAL :",c)
	c =getCode("test@exampl.com")
	fmt.Println("OUTPUT :",c)
}


func getCode(s string)string  {
		h:=hmac.New(sha256.New,[]byte("ourkey"))
		io.WriteString(h,s)
		return fmt.Sprintf("%x", h.Sum(nil))
}