package main

import (
	"errors"
	"testing"
)

type WriterMock struct {
	Buffer     *[]byte
	ThrowError *bool
}

func (w WriterMock) WriteString(string) (int, error) {
	return -1, nil
}

func (w WriterMock) Flush() error {
	return nil
}

type ReaderMock struct {
	Buffer     *[]byte
	ThrowError *bool
}

func (r ReaderMock) ReadBytes(byte) ([]byte, error) {
	return nil, nil
}

type ConnMock struct {
	ReadBuffer *[]byte
	ThrowError *bool
}

func (c ConnMock) Close() error {
	if *c.ThrowError {
		return errors.New("Dummy disconnect error")
	} else {
		c.ReadBuffer = nil
	}
	return nil
}

func (c ConnMock) Read([]byte) (int, error) {
	return -1, nil
}

func (c ConnMock) Write([]byte) (int, error) {
	return -1, nil
}

func TestSetConnection(t *testing.T) {
	var c Client
	golden_host := "localhost"
	golden_port := 8080
	c.SetConnection(golden_host, golden_port)
	host := c.ServerHost
	port := c.ServerPort

	if host != golden_host {
		t.Errorf("(expected) %s != (actual) %s", golden_host, host)
	}

	if port != golden_port {
		t.Errorf("(expected) %d != (actual) %d", golden_port, port)
	}
}

func TestDisconnect(t *testing.T) {
	var b bool
	c := Client{"", -1, ConnMock{nil, &b}, nil, nil}

	b = false
	if !c.Disconnect() {
		t.Error("Failed disconnect unexpected")
	}

	b = true
	if c.Disconnect() {
		t.Error("Successful disconnect unexpected")
	}
}

func TestSendMessage(t *testing.T) {

}

func TestGetResponse(t *testing.T) {

}

/*
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
*/
