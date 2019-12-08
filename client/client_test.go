package main

import (
    "testing"
)

func TestSetConnection(t * testing.T) {
	var c Client
	golden_host := "localhost"
	golden_port := 8080
	c.SetConnection(golden_host, golden_port)
	host := c.ServerHost
	port := c.ServerPort

	if(host != golden_host) {
		t.Errorf("(expected) %s != (actual) %s", golden_host, host)
	}

	if(port != golden_port) {
		t.Errorf("(expected) %d != (actual) %d", golden_port, port)
	}
}


