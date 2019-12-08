package message

import (
    "fmt"
)

const PING_SIZE = 0
const GET_SIZE = 1
const PUT_SIZE = 2

func getSizeError(p Payload) string {
    return fmt.Sprintf("%T", p)
}

type Payload interface {}

type Ping struct {}

type Get  struct {
    Query string
}

type Put  struct {
    Key   string
    Value string
}

/*
func (p Ping) setContents(payload* Payload) {
    return p;
}

func (g Get) setContents(payload* Payload) {
    return g;
}

func (p Put) setContents(payload* Payload) {
    return p;
}
*/