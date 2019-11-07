package message

import (
    "encoding/json"
    "log"
)

type Message struct {
        Head    Header
        Body    string
}

type Header int
const (
    BAD  Header = iota
    PING
    GET
    PUT
)
func (header Header) String() string {
    headers := [...]string{
               "BAD_MESSAGE",
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

func Decode (raw_msg []byte) (Message, bool) {
    var decoded Message
    err := json.Unmarshal(raw_msg, &decoded)
    if(err != nil) {
        return Message{BAD, ""}, false
    }
    return decoded, decoded.Valid()
}

func (msg Message) Encode () []byte {
    encoded, err := json.Marshal(msg)
    if(err != nil) {
        // If encoding craps out thats really bad
        log.Fatal(err)
    }
    return encoded
}
