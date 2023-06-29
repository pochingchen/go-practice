package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

type HashReader interface {
	io.Reader
	hash() string
}

func main() {
	payload := []byte("hello high value software engineer")
	//hashAndBroadcast(bytes.NewReader(payload))
	hashAndBroadcast(NewHashReader(payload))
}

type hashReader struct {
	*bytes.Reader
	buf *bytes.Buffer
}

func NewHashReader(b []byte) *hashReader {
	return &hashReader{
		Reader: bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}

func (h *hashReader) hash() string {
	return hex.EncodeToString(h.buf.Bytes())
}

func hashAndBroadcast(r HashReader) error {
	//b, err := ioutil.ReadAll(r)
	//if err != nil {
	//	return err
	//}
	//
	//hash := sha1.Sum(b)
	//fmt.Println(hex.EncodeToString(hash[:]))

	//hash := r.(*hashReader).hash()
	hash := r.hash()
	fmt.Println(hash)

	return broadcast(r)
}

func broadcast(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	fmt.Println("string of the bytes: ", string(b))

	return nil
}
