package notifications

import (
	"context"
	"sync"
	cid "github.com/ipfs/go-cid"
	pubsub "github.com/btc/pubsub"
	blocks "github.com/ipfs/go-block-format"
)

const bufferSize = 16

type PubSub interface{
	Publish(block blocks.Block)
	Subscribe(ctx context.Context, keys ...*cid.Cid) <- chan blocks.Block
	Shutdown()
}

func New() PubSub {
	return &impl{
		wrapped: *pubsub.New(bufferSize),
		cancel:  make(chan struct{}),
	}
}

type impl struct{
	wrapped  pubsub.PubSub
	// These two fields make up a shutdown "lock"
	// We need them as calling, eg  
	// blocks forever and fixing this in pubsub would be rather invasive

	cancel chan struct{}
	wg sync.WaitGroup
}

func (ps *impl) Publish(block blocks.Block){
	ps.wg.Add(1)
	defer ps.wg.Done()

	select {
	case <- ps.cancel:
		// Already shutdown, bail
		return
	default:
	}
	ps.wrapped.Pub(block, block.Cid().KeyString())
}

// Not safe to call more than once
func (ps *impl) Shutdown(){
	// Interrupt in-progress subscriptions 
	close(ps.cancel)
	// Wait for them to finish
	ps.wg.Wait()
	// shutdown the pubsub
	ps.wrapped.Shutdown()
}
