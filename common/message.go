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
    PING Header = iota
    GET
    PUT
)
func (header Header) String() string {
    headers := [...]string{
               "PING",
               "GET",
               "PUT"}

    if(!header.Valid()) {
        return "UNDEFINED"
    }

    return headers[header]
}
func (header Header) Valid() bool {
    return header >= PING && header <= PUT
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
