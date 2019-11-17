package message

import ()


type Payload interface {
    getContents() []string
    setContents([]string)
}

type Ping struct {}

type Get  struct {
    Query string
}

func (g Get) getContents() []string {
    contents := make([]string, 1)
    contents[0] = g.Query
    return contents
}

type Put  struct {
    Key   string
    Value string
}

