package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	log.Printf("Server initalized")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	conn_handler := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	byt, err := conn_handler.Reader.ReadBytes('\n')
	log.Printf("Error on read: %s", err)
	log.Printf("Recieving: %s", string(byt))
	_, err = conn_handler.Writer.WriteString(string(byt) + "\n")
	log.Printf("Error after writer: %s", err)
	err = conn_handler.Writer.Flush()
	log.Printf("Error after flush: %s", err)
}
