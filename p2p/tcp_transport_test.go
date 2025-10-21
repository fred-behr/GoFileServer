package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":4000"
	tr := NewTCPTransport(listenAddr)

	// listenAddr is the same after initialization
	assert.Equal(t, tr.listenAddress, listenAddr)

	// server
	assert.Nil(t, tr.ListenAndAccept())
}
