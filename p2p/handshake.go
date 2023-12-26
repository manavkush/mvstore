package p2p

// HandshakeFunc
type HandshakeFunc func(Peer) error

// NOPHandshakeFunc is function for protocols where no handshake is required
func NOPHandshakeFunc(Peer) error { return nil }
