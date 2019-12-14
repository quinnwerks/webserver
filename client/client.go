package main

import (
	"bufio"
	"fmt"
	"github.com/quinnwerks/webserver/message"
	"log"
	"net"
	"time"
)

type Writer interface {
	WriteString(string) (int, error)
	Flush() error
}

type Reader interface {
	ReadBytes(byte) ([]byte, error)
}

type Socket interface {
	Close() error
	Read([]byte) (int, error)
	Write([]byte) (int, error)
}

type Client struct {
	ServerHost string
	ServerPort int
	Socket     Socket
	Reader     Reader
	Writer     Writer
}

func (c *Client) SetIO(socket Socket) {
	c.Reader = bufio.NewReader(c.Socket)
	c.Writer = bufio.NewWriter(c.Socket)
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
		c.Reader = nil
		c.Writer = nil
	} else {
		c.Socket = conn
		c.SetIO(c.Socket)
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

func (c *Client) SendMessage(msg message.Message) bool {
	byt := msg.Encode()
	num_bytes, net_err := c.Writer.WriteString(string(byt) + "\n")
	log.Printf("Sent %d Bytes. Original message + newline was %d Bytes.", num_bytes, len(string(byt))+1)
	if net_err != nil {
		log.Printf("Error after writer: %s", net_err)
		return false
	}
	io_err := c.Writer.Flush()
	if io_err != nil {
		log.Printf("Error after flush: %s", io_err)
		return false
	}
	return true
}

func (c *Client) GetResponse() (message.Message, bool) {
	byt, io_err := c.Reader.ReadBytes('\n')
	if io_err != nil {
		log.Printf("Error after read %s", io_err)
		return message.Message{},false
	} else {
		log.Printf("Recieving: %s", string(byt))
		msg := message.Decode(byt)
		return msg, true
	}
}

func main() {
	get_msg := message.Message{Head: message.GET, Body: message.Get{Query: "Hello World"}}
	c := Client{}
	c.SetConnection("localhost", 8080)
	c.Connect()
	c.SendMessage(get_msg)
	time.Sleep(time.Second)
	c.GetResponse()
	c.Disconnect()
}
