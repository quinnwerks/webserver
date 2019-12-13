package main

import (
	"bufio"
	"fmt"
	"github.com/quinnwerks/webserver/message"
	"log"
	"net"
	"time"
)

type Socket interface {
	Close() error
	Read([]byte) (int, error)
	Write([]byte) (int, error)
}

type Client struct {
	ServerHost string
	ServerPort int
	Socket     Socket
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
	conn_str := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial(conn_type, conn_str)
	if err != nil {
		log.Printf("Error in making a connection:%s", err)
		c.Socket = nil
		c.ReadWriter = nil
	} else {
		c.Socket = conn
		c.ReadWriter = bufio.NewReadWriter(bufio.NewReader(c.Socket), bufio.NewWriter(c.Socket))
	}

	return err
}

func (c *Client) Disconnect() bool {
	net_err := c.Socket.Close()
	if net_err != nil {
		log.Printf("Error attempting to disconnect: %s", net_err)
		return false
	}
	return true
}

func (c *Client) SendToServer(msg message.Message) bool {
	byt := msg.Encode()
	num_bytes, net_err := c.ReadWriter.Writer.WriteString(string(byt) + "\n")
	log.Printf("Sent %d Bytes. Original message + newline was %d Bytes.", num_bytes, len(string(byt))+1)
	if net_err != nil {
		log.Printf("Error after writer: %s", net_err)
	}
	io_err := c.ReadWriter.Flush()
	if io_err != nil {
		log.Printf("Error after flush: %s", io_err)
	}
	return true
}

func (c *Client) GetResponse() message.Message {
	byt, io_err := c.ReadWriter.Reader.ReadBytes('\n')
	if io_err != nil {
		log.Printf("Error after read %s", io_err)
	}
	log.Printf("Recieving: %s", string(byt))
	msg := message.Decode(byt)
	return msg
}

func main() {
	get_msg := message.Message{Head: message.GET, Body: message.Get{Query: "Hello World"}}
	c := Client{}
	c.SetConnection("localhost", 8080)
	c.Connect()
	c.SendToServer(get_msg)
	time.Sleep(time.Second)
	c.GetResponse()
	c.Disconnect()
}
