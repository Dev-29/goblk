package network

import (
	"fmt"
	"sync"
)

// struct representing a local transport
type LocalTransport struct {
	addr      NetAddr                     // address of the local transport
	consumeCh chan RPC                    // channel to consume messages from
	lock      sync.RWMutex                // lock to protect the peers map
	peers     map[NetAddr]*LocalTransport // map of peers
}

// create a new local transport
func NewLocalTransport(addr NetAddr) *LocalTransport {
	return &LocalTransport{
		addr:      addr,                              // address of the local transport
		consumeCh: make(chan RPC, 1024),              // channel to consume messages from
		peers:     make(map[NetAddr]*LocalTransport), // map of peers
	}
}

// returns a channel to consume messages from
func (t *LocalTransport) Consume() <-chan RPC {
	return t.consumeCh
}

// connects to a remote transport
func (t *LocalTransport) Connect(tr *LocalTransport) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.peers[tr.Addr()] = tr

	return nil
}

// sends a message to a remote transport
func (t *LocalTransport) SendMessage(to NetAddr, payload []byte) error {
	t.lock.RLock()
	defer t.lock.Unlock()

	// check if the peer exists
	peer, ok := t.peers[to]
	if !ok {
		return fmt.Errorf("%s: could not send message to %s", t.addr, to)
	}

	// send the message to the peer
	peer.consumeCh <- RPC{
		From:    t.addr,
		Payload: payload,
	}

	return nil
}

// returns the address of the local transport
func (t *LocalTransport) Addr() NetAddr {
	return t.addr
}
