package message

import (
    "testing"
)

func GoodHeaderCreate(t *testing.T, header_int int) {
    golden_header :=  Header(header_int)
    golden_body   := "never eat shredded wheat"

    msg := Create(golden_header, golden_body)

    if(!msg.header.Valid()) {
        t.Error("invalid message")
    }
    if(msg.header != golden_header) {
        t.Error("golden != result", golden_header, msg.header)
    }
    if(msg.body != golden_body) {
        t.Error("golden != result:", golden_body, msg.body)
    }
}

func BadHeaderCreate(t *testing.T, header_int int){
    golden_header := Header(header_int)
    golden_body   := "this message is bad"

    msg := Create(golden_header, golden_body)
    header := msg.header

    if(header.Valid()) {
        t.Error("valid message")
    }
    if(header.String() != "UNDEFINED") {
        t.Error("golden != result", "UNDEFINED", msg.header)
    }
    if(msg.body != golden_body) {
        t.Error("golden != result:", golden_body, msg.body)
    }

}

func TestMessageRanges(t *testing.T){
    t.Run("HeaderLowRangeInvalid", func (t *testing.T) {
        value := -1
        BadHeaderCreate(t, value)
    })

    t.Run("HeaderHighRangeInvalid", func (t *testing.T) {
        value := 3
        BadHeaderCreate(t, value)
    })

    t.Run("HeaderInRange", func (t * testing.T) {
        value := 2
        GoodHeaderCreate(t, value)
    })
}

