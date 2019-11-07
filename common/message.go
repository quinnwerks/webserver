package message

import (
    "encoding/json"
    "fmt"
    "log"
)

type Message struct {
        Head    Header
        Body    string
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

func (msg Message) Valid() bool {
    return msg.Head.Valid()
}

func Create (header Header, body string) Message {
    msg := Message{header, body}
    return msg
}

func (msg Message)String() string {
    ret := "header:" + msg.Head.String() + "\n" +
           "body:" +   msg.Body            + "\n"
    return ret
}

func Decode (raw_msg []byte) (Message, bool) {
    var decoded Message
    err := json.Unmarshal(raw_msg, &decoded)
    if(err != nil) {
        fmt.Println("todo decode")
    }
    fmt.Println("Decoded", decoded.Head)
    return decoded, decoded.Valid()
}

func (msg Message) Encode () []byte {
    encoded, err := json.Marshal(msg)
    if(err != nil) {
        log.Fatal(err)
    }
    return encoded
}
