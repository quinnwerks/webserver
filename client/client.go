package main

import (
    "fmt"
    "net"
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

func main() {
    conn, err := net.Dial("tcp", ":8080")
    if err != nil {
        fmt.Println("Handle dial error")
    }

    fmt.Fprintf(conn, "Hello world!!!")
}
