package message

import (
    "testing"
)
func TestDecodeEncodeGet(t * testing.T) {
    var get *Get;
    str := "blah"
    msg := Message{GET, Get{str}}
    byt := msg.Encode()
    new_msg := Decode(byt)
 
    get = new_msg.Body.(*Get)
    if(get.Query != str) {
        t.Error("query != golden")
    }
}
func TestDecodeEncodePut(t * testing.T) {
    var put *Put;
    key := "blah"
    val := "blub"
    msg := Message{PUT, Put{key, val}}
    byt := msg.Encode()
    new_msg := Decode(byt)
 
    put = new_msg.Body.(*Put)
    if(put.Key != key) {
        t.Error("key != golden")
    }
    if(put.Value != val) {
        t.Error("val != golden")
    }
}
