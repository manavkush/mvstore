package p2p

// HandshakeFunc
type HandshakeFunc func(any) error

// NOPHandshakeFunc is function for protocols where no handshake is required
func NOPHandshakeFunc(any) error { return nil }
