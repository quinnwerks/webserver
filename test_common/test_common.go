package test_common

import (
	"errors"
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
