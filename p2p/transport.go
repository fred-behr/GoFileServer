package p2p

// Peer is an interface, that represents remote nodes
type Peer interface {
}

// Transport is anything that handles communication
// between nodes in the network. This can be in the
// form of TCP, UDP, Web Sockets etc.
type Transport interface {
	ListenAndAccept() error
}
