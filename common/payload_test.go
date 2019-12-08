package message

import (
    "testing"
)

func TestGet (t *testing.T) {
    const GOLDEN_QUERY   = "Blah"
    const GOLDEN_QUERY_2 = "New"

    test := Get{"Blah"}
    if(test.Query != GOLDEN_QUERY) {
        t.Error("golden != actual", GOLDEN_QUERY, test.Query)
    }
}

func TestPut (t *testing.T) {
    const GOLDEN_KEY     = "Key"
    const GOLDEN_VALUE   = "To Success"

    test := Put{GOLDEN_KEY, GOLDEN_VALUE}

    if(test.Key != GOLDEN_KEY) {
        t.Error("golden != actual", GOLDEN_KEY, test.Key)
    }
    if(test.Value != GOLDEN_VALUE) {
        t.Error("golden != actual", GOLDEN_VALUE, test.Value)
    }
}
