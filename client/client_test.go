package main

import (
	"testing"
	"github.com/quinnwerks/webserver/message"
	"github.com/quinnwerks/webserver/test_common"
)

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
	var should_flush, io_err, flush_err, net_err bool
	c := MakeClientForTests(&should_flush, &io_err, &flush_err, &net_err)

	net_err = false
	if !c.Disconnect() {
		t.Error("Failed disconnect unexpected")
	}

	net_err = true
	if c.Disconnect() {
		t.Error("Successful disconnect unexpected")
	}
}

func MakeClientForTests(should_flush, io_err, flush_err, net_err *bool) Client {
	buff    := make([]byte,0)
	conn   := test_common.ConnMock{Buffer:&buff,
					               ThrowError:net_err}
	writer := test_common.WriterMock {Conn:&conn,
					      ThrowIOErr:io_err,
					      ThrowFlushErr:flush_err,
						  ShouldFlush:should_flush}
	
	reader := test_common.ReaderMock {Conn:&conn,
						  ThrowIOErr:io_err}

	client := Client{ServerHost:"", 
				     ServerPort:-1, 
				     Socket:conn, 
				     Reader:reader, 
					 Writer:writer}
					 
	return client
}

func TestSendMessage(t *testing.T) {
	var should_flush, io_err, flush_err, net_err bool
	get_msg := message.Message{
								Head: message.GET, 
								Body: 
								message.Get {
								Query: "Hello World"}}
	golden_buff := append(get_msg.Encode(), '\n')
	c := MakeClientForTests(&should_flush, &io_err, &flush_err, &net_err)
	c.SendMessage(get_msg)
	buff,_ := c.Reader.ReadBytes(' ')	

	if string(buff) != string(golden_buff) {
		t.Errorf("Message not same as original")
	}

	io_err = true
	should_flush = true
	if c.SendMessage(get_msg) {
		t.Errorf("Write msg succeeded when failure expected")
	}

	io_err = false
	flush_err = true
	if c.SendMessage(get_msg) {
		t.Errorf("Flush msg succeeded when failure expected")
	}

}

func TestGetResponse(t *testing.T) {
	var should_flush, io_err, flush_err, net_err bool
	get_msg := message.Message{
		Head: message.GET, 
		Body: 
		message.Get {
		Query: "Hello World"}}
	c := MakeClientForTests(&should_flush, &io_err, &flush_err, &net_err)
	c.Writer.WriteString(string(get_msg.Encode()) + "\n")
	rsp,_ := c.GetResponse()

	if rsp.Head != message.GET {
		t.Errorf("Message is of wrong type")
	}

	io_err = true
	_,is_error := c.GetResponse()
	if is_error {
		t.Errorf("Read message suceeded when failure expected")
	}

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
