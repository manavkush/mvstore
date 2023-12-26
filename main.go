package main

import (
	"fmt"
	"log"

	"github.com/manavkush/mvstore/p2p"
)

func OnPeerConnected(p p2p.Peer) error {
	fmt.Println("Doing some logic outside the TCPTransport")
	return nil
}

func main() {
	tcpConfig := p2p.TCPTransportConfig{
		ListenAddr:      ":4000",
		HandshakeFunc:   p2p.NOPHandshakeFunc,
		Decoder:         p2p.DefaultDecoder{},
		OnPeerConnected: OnPeerConnected,
	}
	tr := p2p.NewTCPTransport(tcpConfig)

	go func() error {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%+v\n", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
