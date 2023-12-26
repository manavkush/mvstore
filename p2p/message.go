package p2p

import "net"

// RPC holds any data that's being sent over
// every transport between two peers in the network
type RPC struct {
	From    net.Addr
	Payload []byte
}
