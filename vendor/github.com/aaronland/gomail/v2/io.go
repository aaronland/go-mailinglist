package gomail

import (
	"bytes"
	"io"
)

type ReadSeekCloser interface {
	io.Reader
	io.Seeker
	io.Closer
}

type readSeekCloser struct {
	ReadSeekCloser
	reader *bytes.Reader
}

func (r *readSeekCloser) Read(p []byte) (int, error) {
	return r.reader.Read(p)
}

func (r *readSeekCloser) Seek(offset int64, whence int) (int64, error) {
	return r.reader.Seek(offset, whence)
}

func (r *readSeekCloser) Close() error {
	return nil
}

func NewReadSeekCloserFromBuffer(buf bytes.Buffer) ReadSeekCloser {

	r := bytes.NewReader(buf.Bytes())
	return NewReadSeekCloser(r)
}

func NewReadSeekCloser(r *bytes.Reader) ReadSeekCloser {

	rsc := readSeekCloser{
		reader: r,
	}

	return &rsc
}
