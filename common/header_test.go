package message

import (
    "testing"
)

func GoodHeaderCreate(t *testing.T, header_int int) {
    header :=  Header(header_int)

    if(!header.Valid()) {
        t.Error("invalid message")
    }
}

func BadHeaderCreate(t *testing.T, header_int int){
    header := Header(header_int)
    golden_str := "Bad"

    if(header.Valid()) {
        t.Error("valid message")
    }
    if(string(header) != golden_str) {
        t.Error("header != Bad")
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
