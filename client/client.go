package main

import (
    "fmt"
    "net"
//    "bufio"
)

func main() {
    conn, err := net.Dial("tcp", ":8080")
    if err != nil {
        fmt.Println("Handle dial error")
    }

    fmt.Fprintf(conn, "Hello world!!!")
}
