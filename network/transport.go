package network

type NetAddr string

// RPC is a struct representing a remote procedure call.
type RPC struct {
	From    NetAddr
	Payload []byte
}

// Transport is an interface that represents a network transport.
type Transport interface {
	Consume() <-chan RPC               // returns a channel to consume messages from
	Connect(Transport) error           // connects to a remote transport
	SendMessage(NetAddr, []byte) error // sends a message to a remote transport
	Addr() NetAddr                     // returns the address of the transport
}
