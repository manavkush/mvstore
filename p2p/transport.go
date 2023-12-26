package p2p

// Peer is the interface that represents a peer in the network
type Peer interface {
	Close() error
}

// Transport is the interface for the transport layer
// It is responsible for sending and receiving messages between peers
// This can be of form (TCP, UDP, websockets, etc).
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
