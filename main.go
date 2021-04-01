package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	var (
		target     string
		port       int
		serverPort string
	)
	target = "192.168.99.103"
	port = 3000
	serverPort = ":8000"

	listener, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handle(conn, target, port)
	}
}

func handle(src net.Conn, target string, port int) {
	address := fmt.Sprintf("%s:%d", target, port)
	dst, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalln("Unable to connect to the unreactable host")
	}
	defer dst.Close()

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}
