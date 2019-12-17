package server

import (
	"bufio"
	"log"
	"net"
)

func main() {
	log.Printf("Server initalized")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			// handle error
		} else {
			go handleConnection(conn)
		}
	}
}

func handleConnection(conn net.Conn) {
	connHandler := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	byt, err := connHandler.Reader.ReadBytes('\n')
	log.Printf("Error on read: %s", err)
	log.Printf("Recieving: %s", string(byt))
	_, err = connHandler.Writer.WriteString(string(byt) + "\n")
	log.Printf("Error after writer: %s", err)
	err = connHandler.Writer.Flush()
	log.Printf("Error after flush: %s", err)
}
