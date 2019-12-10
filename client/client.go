package main

import (
	"bufio"
	"github.com/quinnwerks/webserver/message"
	"log"
	"net"
	"time"
)

type Client struct {
	ServerHost string
	ServerPort int
	Socket     net.Conn
	ReadWriter *bufio.ReadWriter
}

func (c *Client) SetConnection(host string, port int) {
	c.ServerHost = host
	c.ServerPort = port
}

func (c *Client) Connect() error {
	conn_type := "tcp"
	host := c.ServerHost
	port := c.ServerPort
	conn_str := host + ":" + string(port)
	conn, err := net.Dial(conn_type, conn_str)
	if err != nil {
		log.Printf("%s", err)
	} else {
		c.Socket = conn
	}

	return err
}

func (c *Client) HandleConnectionError(err error) bool {
	return false
}

func main() {
	get_msg := message.Message{Head: message.GET, Body: message.Get{Query: "Hello World"}}
	byt := get_msg.Encode()
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Println("Handle dial error")
	}
	log.Printf("Sending: %s", string(byt))

	conn_handler := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	_, err = conn_handler.Writer.WriteString(string(byt) + "\n")
	log.Printf("Error after writer: %s", err)
	err = conn_handler.Writer.Flush()
	log.Printf("Error after flush: %s", err)
	time.Sleep(time.Second)
	byt, err = conn_handler.Reader.ReadBytes('\n')
	log.Printf("Error after read %s", err)
	log.Printf("Recieving: %s", string(byt))
	conn.Close()
}
