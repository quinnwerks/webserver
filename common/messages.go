package messages

import (
    "encoding/json"
    "fmt"
)

type Message struct {
        header  string
        body    string
}

func Decode (raw_msg []byte) Message {
    decoded := Message{}
    err := json.Unmarshal(raw_msg, &decoded)
    if(err != nil){
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
