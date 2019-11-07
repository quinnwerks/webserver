package message

import (
    "testing"
)

func GoodHeaderCreate(t *testing.T, header_int int) {
    golden_header :=  Header(header_int)
    golden_body   := "never eat shredded wheat"

    msg := Create(golden_header, golden_body)

    if(!msg.Head.Valid()) {
        t.Error("invalid message")
    }
    if(msg.Head != golden_header) {
        t.Error("golden != result", golden_header, msg.Head)
    }
    if(msg.Body != golden_body) {
        t.Error("golden != result:", golden_body, msg.Body)
    }
}

func BadHeaderCreate(t *testing.T, header_int int){
    golden_header := Header(header_int)
    golden_body   := "this message is bad"

    msg := Create(golden_header, golden_body)
    header := msg.Head

    if(header.Valid()) {
        t.Error("valid message")
    }
    if(header.String() != "UNDEFINED") {
        t.Error("golden != result", "UNDEFINED", msg.Head)
    }
    if(msg.Body != golden_body) {
        t.Error("golden != result:", golden_body, msg.Body)
    }

}

func TestMessageRanges(t *testing.T){
    t.Run("HeaderLowRangeInvalid", func (t *testing.T) {
        value := -1
        BadHeaderCreate(t, value)
    })

    t.Run("HeaderHighRangeInvalid", func (t *testing.T) {
        value := 4
        BadHeaderCreate(t, value)
    })

    t.Run("HeaderInRange", func (t * testing.T) {
        value := 2
        GoodHeaderCreate(t, value)
    })
}

func TestEncodeGood(t *testing.T) {
    msg := Create(GET, "hello")
    enc_msg := string(msg.Encode());
    golden_enc_msg := "{\"Head\":2,\"Body\":\"hello\"}"
    if(enc_msg != golden_enc_msg){
        t.Error("golden != result:", golden_enc_msg, enc_msg)
    }
}

func TestDecodeGood(t * testing.T) {
    enc_msg := []byte("{\"Head\":3,\"Body\":\"hello\"}")
    golden_msg := Create(PUT, "hello")
    msg, valid := Decode(enc_msg)
    if(msg != golden_msg || !valid) {
        t.Error("golden != result:", golden_msg, msg)
    }
}
