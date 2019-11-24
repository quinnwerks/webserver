package message

import (
    "testing"
)

func TestGet (t *testing.T) {
    const GOLDEN_QUERY   = "Blah"
    const GOLDEN_QUERY_2 = "New"

    test := Get{"Blah"}
    contents := test.getContents()
    if(len(contents) != GET_SIZE) {
        t.Error("golden != actual", GET_SIZE, len(contents))
    }
    if(contents[0] != GOLDEN_QUERY) {
        t.Error("golden != actual", GOLDEN_QUERY, contents[0])
    }

    new_contents := make([]string, 1)
    new_contents[0] = GOLDEN_QUERY_2
    test.setContents(new_contents)
    if(test.Query != GOLDEN_QUERY_2) {
        t.Error("golden != actual", GOLDEN_QUERY_2, test.Query)
    }
}

func TestPut (t *testing.T) {
    const GOLDEN_KEY     = "Key"
    const GOLDEN_VALUE   = "To Success"
    const GOLDEN_KEY_2   = "Is To"
    const GOLDEN_VALUE_2 = "Never Give Up"

    test := Put{GOLDEN_KEY, GOLDEN_VALUE}
    contents := test.getContents()
    if(len(contents) != PUT_SIZE) {
        t.Error("golden != actual", PUT_SIZE, len(contents))
    }
    if(contents[0] != GOLDEN_KEY) {
        t.Error("golden != actual", GOLDEN_KEY, contents[0])
    }
    if(contents[1] != GOLDEN_VALUE) {
        t.Error("golden != actual", GOLDEN_VALUE, contents[1])
    }

    new_contents := make([]string, 2)
    new_contents[0] = GOLDEN_KEY_2
    new_contents[1] = GOLDEN_VALUE_2
    test.setContents(new_contents)
    if(test.Key != GOLDEN_KEY_2) {
        t.Error("golden != actual", GOLDEN_KEY_2, test.Key)
    }
    if(test.Value != GOLDEN_VALUE_2) {
        t.Error("golden != actual", GOLDEN_VALUE_2, test.Value)
    }
}
