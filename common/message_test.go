package message

import (
    "testing"
    "fmt"
)
/*
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
*/
func TestDecodePayload(t * testing.T) {
    msg := Message{GET, Get{"blah"}}
    byt := msg.Encode()
    new_msg := Decode(byt)
    fmt.Println(new_msg)
    var get Get;
    GetType(new_msg, &get) 
    fmt.Println(new_msg.Body.Query)
}
