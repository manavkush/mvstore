package p2p

import "net"

// Message holds any data that's being sent over
// every transport between two peers in the network
type Message struct {
	From    net.Addr
	Payload []byte
}
