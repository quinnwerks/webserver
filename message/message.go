package message

import (
    "encoding/json"
    "log"
    "fmt"
)

type RawMessage struct {
    Head Header
    Body json.RawMessage
}

type Message struct {
        Head Header
        Body Payload
}

func Decode (raw_json []byte) Message {
    var raw_msg RawMessage
    err := json.Unmarshal(raw_json, &raw_msg)
    if(err != nil) {
        log.Fatalln("Error:", err)
        fmt.Println(err)
    }

    var msg Message
    var payload Payload
    switch raw_msg.Head {
        case PING: payload = new(Ping)
        case GET:  payload = new(Get)
        case PUT:  payload = new(Put)
        default:   panic("Message not implemented")
    }

    err = json.Unmarshal(raw_msg.Body, payload)
    if(err != nil) {
        log.Fatalln("Error:",err)
        fmt.Println(err)
    }
    msg.Head = raw_msg.Head
    msg.Body = payload
    return msg
}

func (msg Message) Encode () []byte {
    encoded, err := json.Marshal(msg)
    if(err != nil) {
        // If encoding craps out thats really bad
        fmt.Println(err)
        log.Fatal(err)
    }
    return encoded
}
