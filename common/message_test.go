package message

import (
    "testing"
)

func CreateMessage(t *testing.T) {
    golden_header := "common_phrase"
    golden_body   := "never eat shredded wheat"

    msg := Create(golden_header, golden_body)

    if(msg.header != golden_header) {
        t.Error("golden != result", golden_header, msg.header)
    }
    if(msg.body != golden_body) {
        t.Error("golden != result:", golden_body, msg.body)
    }
}
