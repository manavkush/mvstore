package p2p

import (
	"encoding/gob"
	"io"
)

type Decoder interface {
	Decode(io.Reader, *Message) error
}

type GobDecoder struct{}

func (dec GobDecoder) Decode(r io.Reader, v *Message) error {
	return gob.NewDecoder(r).Decode(v)
}

type DefaultDecoder struct{}

func (dec DefaultDecoder) Decode(r io.Reader, v *Message) error {
	msg := make([]byte, 1024)
	n, err := r.Read(msg)
	if err != nil {
		return err
	}

	v.Payload = msg[:n]
	return nil
}
