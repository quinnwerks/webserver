package main

import (
    "fmt"
    "net"
    "github.com/quinnwerks/webserver/message"
)

func main() {
    conn, err := net.Dial("tcp", ":8080")
    if err != nil {
        fmt.Println("Handle dial error")
    }

    fmt.Fprintf(conn, "Hello world!!!")
}
