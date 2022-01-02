package app

import (
	"fmt"
	"io"
	"log"
	"net"
)

// CreateProxy creates a proxy server
func CreateProxy(src net.Conn, target string, port int) {
	address := fmt.Sprintf("%s:%d", target, port)
	dst, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalln("unable to connect to the unreactable host")
	}
	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
	defer dst.Close()
}
