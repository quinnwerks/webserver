package main

import (
    "log"
    "net"
    "fmt"
    "time"
    "github.com/quinnwerks/webserver/message"
)

type Client struct{
    ServerHost string
    ServerPort int
    Socket     net.Conn 
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
    if(err != nil) {
        log.Printf("%s", err)
    } else {
        c.Socket = conn
    }

    return err
}

func (c * Client) HandleConnectionError(err error) bool{
    return false
}

func main() {
    get_msg := message.Message{message.GET, message.Get{Query:"Hello World"}}
    byt := get_msg.Encode()
    conn, err := net.Dial("tcp", ":8080")
    if err != nil {
        log.Println("Handle dial error")
    }
    fmt.Println("Sending: ", string(byt))
    conn.Write(byt)
    byt = make([]byte, 512)
    time.Sleep(time.Second)
    _, err = conn.Read(byt)
    fmt.Println(err)
    fmt.Println("Recieving: ", string(byt))
    conn.Close()
    fmt.Println("Here")
}
