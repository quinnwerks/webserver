package message

import (
    "encoding/json"
    "fmt"
)

type Message struct {
        header  Header
        body    string
}

type Header int
const (
    PING Header = 0
    GET  Header = 1
    PUT  Header = 2
)
func (header Header) String() string {
    headers := [...]string{
                          "PING",
                          "GET",
                          "PUT"}

    if(header < PING || header > PUT) {
        return "UNDEFINED"
    }

    return headers[header]
}

func Create (header Header, body string) Message {
    msg := Message{header, body}
    return msg
}

func Decode (raw_msg []byte) Message {
    decoded := Message{}
    err := json.Unmarshal(raw_msg, &decoded)
    if(err != nil) {
        fmt.Println("todo decode")
    }
    return decoded
}

func (msg Message) Encode () []byte {
    encoded, err := json.Marshal(msg)
    if(err != nil) {
        fmt.Println("todo encode")
    }
    return encoded
}
