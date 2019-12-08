package main

import (
    "log"
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
    /*
    conn, err := net.Dial("tcp", ":8080")
    if err != nil {
        log.Println("Handle dial error")
    }

    fmt.Fprintf(conn, "Hello world!!!")*/
}
