package message

import (
    "testing"
)

func TestCreate(t *testing.T) {
    golden_header :=  PING
    golden_body   := "never eat shredded wheat"

    msg := Create(golden_header, golden_body)

    if(msg.header != golden_header) {
        t.Error("golden != result", golden_header, msg.header)
    }
    if(msg.body != golden_body) {
        t.Error("golden != result:", golden_body, msg.body)
    }
}

func TestCreateBad(t *testing.T){
    golden_header := Header(-1)
    golden_body   := "this message is bad"

    msg := Create(golden_header, golden_body)
    header := msg.header.String()

    if(header != "UNDEFINED") {
        t.Error("golden != result", golden_header, msg.header)
    }
    if(msg.body != golden_body) {
        t.Error("golden != result:", golden_body, msg.body)
    }

}

