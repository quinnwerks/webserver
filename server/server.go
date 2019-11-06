package main

import (
    "fmt"
    "net"
    "bufio"
)


func main() {
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
    msg, thing := bufio.NewReader(conn).ReadString('\n')
    fmt.Println(conn, msg, thing)
}



