package decision

import (
	"github.com/ipfs/go-block-format"
	"gx/ipfs/QmZoWKhxUmZ2seW4BzX6fJkNR8hh9PsGModr7q171yq2SS/go-libp2p-peer"
	"context"
	"sync"
	"time"
	logging "github.com/ipfs/go-log"
)

var log = logging.Logger("engine")

const (
	// outbox
	outboxChanBuffer = 0
)

// Envelope contains a message for a peer
type Envelope struct{
	// peer is the intended recipient
	Peer peer.ID
	// Block is the payload
	Block blocks.Block
	// A callback to notify decision queue that task is complete
	Sent func()
}

type Engine struct{
	// 
	// peerRequestQueue *prq
}