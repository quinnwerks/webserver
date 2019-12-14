package main

import (
	"errors"
	"testing"
	"github.com/quinnwerks/webserver/message"
)

type WriterMock struct {
	Conn          *ConnMock
	ThrowIOErr    *bool
	ThrowFlushErr *bool
	ShouldFlush   *bool
}

func (w WriterMock) WriteString(s string) (int, error) {
	if *w.ThrowIOErr {
		return -1, errors.New("Dummy write string error")
	}
	return w.Conn.Write([]byte(s))
}

func (w WriterMock) Flush() error {
	if *w.ThrowFlushErr {
		return errors.New("Dummy flush error")
	}
	if *w.ShouldFlush {
		*w.Conn.Buffer = (*w.Conn.Buffer)[:0]
	}
	return nil
}

type ReaderMock struct {
	Conn         *ConnMock
	ThrowIOErr   *bool
}

func (r ReaderMock) ReadBytes(byte) ([]byte, error) {
	if *r.ThrowIOErr {
		return nil, errors.New("Dummy io error")
	}
	return *r.Conn.Buffer, nil
}

type ConnMock struct {
	Buffer *[]byte
	ThrowError *bool
}

func (c ConnMock) Close() error {
	if *c.ThrowError {
		return errors.New("Dummy disconnect error")
	} else {
		c.Buffer = nil
	}
	return nil
}

func (c ConnMock) Read([]byte) (int, error) {
	return -1, nil
}

func (c ConnMock) Write(b []byte) (int, error) {
	num_bytes := 0
	for _, char := range b {
		*c.Buffer = append(*c.Buffer, char)
		num_bytes++;
	}
	return num_bytes, nil
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
	var throw_err bool
	c := Client{"", 
				-1, 
				ConnMock{nil, &throw_err}, 
				nil, 
				nil}

	throw_err = false
	if !c.Disconnect() {
		t.Error("Failed disconnect unexpected")
	}

	throw_err = true
	if c.Disconnect() {
		t.Error("Successful disconnect unexpected")
	}
}

func MakeClientForTests(should_flush, io_err, flush_err *bool) Client {
	buff    := make([]byte,0)
	conn   := ConnMock{Buffer:&buff,
					   ThrowError:nil}
	writer := WriterMock {Conn:&conn,
					      ThrowIOErr:io_err,
					      ThrowFlushErr:flush_err,
						  ShouldFlush:should_flush}
	
	reader := ReaderMock {Conn:&conn,
						  ThrowIOErr:io_err}

	client := Client{ServerHost:"", 
				     ServerPort:-1, 
				     Socket:conn, 
				     Reader:reader, 
					 Writer:writer}
					 
	return client
}

func TestSendMessage(t *testing.T) {
	var should_flush, io_err, flush_err bool
	get_msg := message.Message{
								Head: message.GET, 
								Body: 
								message.Get {
								Query: "Hello World"}}
	golden_buff := append(get_msg.Encode(), '\n')
	c := MakeClientForTests(&should_flush, &io_err, &flush_err)
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
	var should_flush, io_err, flush_err bool
	get_msg := message.Message{
		Head: message.GET, 
		Body: 
		message.Get {
		Query: "Hello World"}}
	c := MakeClientForTests(&should_flush, &io_err, &flush_err)
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
