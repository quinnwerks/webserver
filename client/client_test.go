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

func AssertOnErrorMessage(t * testing.T, err string, golden_msg string) {
	if(err != golden_msg) {
		t.Errorf("(expected) %s != (actual) %s", golden_msg, err)
	}
} 

func ExpectErrorConnect(t * testing.T, host string, port int, msg string) {
	c := Client{ServerHost:"blah", ServerPort:0}
	err := c.Connect().Error()
	golden_err := "dial tcp: lookup " + host  + ": " + msg
	AssertOnErrorMessage(t, err, golden_err)
}

func TestBadConnect(t * testing.T) {
	t.Run("NoSuchHost", func (t * testing.T) {
		host := "blah"
		port := -1
		msg := "no such host"
		ExpectErrorConnect(t, host, port, msg)
	})
}