package message

import (
    "fmt"
    "reflect"
)

const PING_SIZE = 0
const GET_SIZE = 1
const PUT_SIZE = 2

func getSizeError(p Payload) string {
    return fmt.Sprintf("%T", p)
}

type Payload interface {
//    getContents()
}

type PayloadPtr interface {
    setContents(Payload)
}

func GetType(p Payload, i interface{}) {
    type_of_payload := reflect.TypeOf(p)
    fmt.Println(type_of_payload)
}

type Ping struct {}

type Get  struct {
    Query string
}

type Put  struct {
    Key   string
    Value string
}

func (p Ping) getContents() []string {
    contents := make([]string, PING_SIZE)
    return contents
}

func (p *Ping) setContents(s []string) {
    panic(fmt.Sprintf("Can't set contents of a %T", p))
}


func (g Get) getContents() []string {
    contents := make([]string, GET_SIZE)
    contents[0] = g.Query
    return contents
}

func (g *Get) setContents(s []string) {
    if(len(s) != GET_SIZE) {
        panic(getSizeError(g))
    }
    g.Query = s[0]
}

func (p Put) getContents() []string {
    contents := make([]string, PUT_SIZE)
    contents[0] = p.Key
    contents[1] = p.Value
    return contents
}

func (p *Put) setContents(s []string) {
    if (len(s) != PUT_SIZE) {
        panic(getSizeError(p))
    }
    p.Key = s[0]
    p.Value = s[1]
}
