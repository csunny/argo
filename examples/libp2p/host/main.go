/* For the most applications , the host is the basic 
building block you'll need to get started. The guide will show how to
constrant and use a simple host.package main

The host is an abstraction that managers services on top of a swarm,
it provides a clean interface to connect to a service on a given remote peer.

if you want to create a host with a default configuration, you can do the following:
*/

package main

import (
	"context"
	"crypto/rand"
	"fmt"

	libp2p "github.com/libp2p/go-libp2p"
	crypto "github.com/libp2p/go-libp2p-crypto"
)

func main(){
// The context governs the lifetime of the libp2p node

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// to construct a simple host with all the default settings, just use New
	h, err := libp2p.New(ctx)
	if err != nil{
		panic(err)
	}

	fmt.Printf("Hello world, my hosts ID is %s\n", h.ID())

	// set your own keypair
	priv, _, err := crypto.GenerateEd25519Key(rand.Reader) 
	if err != nil{
		panic(err)
	}

	h2, err := libp2p.New(ctx, libp2p.Identity(priv), 
		// set your owner listen address
		// the config takes an array of addresses, specify as many as
		// you want.
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/9000"),
	)
	if err != nil{
		panic(err)
	}

	fmt.Printf("Hello world, my second hosts ID is %s\n", h2.ID())
}