package main

import (
	"log"

	"github.com/manavkush/mvstore/p2p"
)

func main() {
	tcpConfig := p2p.TCPTransportConfig{
		ListenAddr:    ":4000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpConfig)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
