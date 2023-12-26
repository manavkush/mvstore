package p2p

import (
	"fmt"
	"net"
)

// TCPPeer represents a remote node over a TCP established connection
type TCPPeer struct {
	// conn is the underlying connection of the peer
	conn net.Conn

	// If we dial and retrieve a conn => outbound == true
	// If we accept and retrieve a conn => outbound == false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	TCPTransportConfig
	listener net.Listener
	rpcCh    chan RPC
}

type TCPTransportConfig struct {
	ListenAddr      string
	HandshakeFunc   HandshakeFunc
	Decoder         Decoder
	OnPeerConnected func(peer Peer) error
}

func NewTCPTransport(config TCPTransportConfig) *TCPTransport {
	return &TCPTransport{
		TCPTransportConfig: config,
		rpcCh:              make(chan RPC),
	}
}

// Close implements the Peer interface
// Closes the connection
func (p *TCPPeer) Close() error {
	return p.conn.Close()
}

// Consume implements the Transport interface
// It returns a channel of RPCs that are received over the transport
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcCh
}

func (t *TCPTransport) ListenAndAccept() error {
	ln, err := net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}

	t.listener = ln

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}

		go t.handleConn(conn)
	}
}

type Temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn) {
	var err error
	defer func() {
		fmt.Printf("Closing peer connection: %v\n", err)
		conn.Close()
	}()

	// Create a new peer
	tcpPeer := NewTCPPeer(conn, true)

	// Handshake
	if err = t.HandshakeFunc(tcpPeer); err != nil {
		return
	}

	// Do something with the peer
	if t.OnPeerConnected != nil {
		if err = t.OnPeerConnected(tcpPeer); err != nil {
			return
		}
	}

	// Read Loop
	rpc := RPC{}

	fmt.Println("Starting Reading print loop")
	for {
		err := t.Decoder.Decode(conn, &rpc)

		if err == net.ErrClosed {

		} else if err != nil {
			fmt.Printf("TCP error: %s\n", err)
			continue
		}

		rpc.From = conn.RemoteAddr()

		t.rpcCh <- rpc
		// n, err := conn.Read(buff)
	}

}
