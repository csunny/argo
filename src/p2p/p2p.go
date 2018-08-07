package p2p

import (
	"encoding/hex"
	"time"
	"crypto/sha256"
	"sync"
	"context"
	"fmt"
	"log"
	"crypto/rand"
	"io"
	"bufio"
	"encoding/json"
	"os"
	"strings"
	"flag"
	"strconv"

	mrand "math/rand"

	"github.com/libp2p/go-libp2p-crypto"
	"github.com/libp2p/go-libp2p-host"

	"github.com/libp2p/go-libp2p"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/libp2p/go-libp2p-net"

	"github.com/davecgh/go-spew/spew"
	golog "github.com/ipfs/go-log"
	gologging "github.com/whyrusleeping/go-logging"

	"github.com/libp2p/go-libp2p-peer"

	pstore "github.com/libp2p/go-libp2p-peerstore"
)

var Blockchain []Block

var mutex = &sync.Mutex{}

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

// sha256 hashing
func calculateHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}

func calculateBlockHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash

	return calculateHash(record)
}

func generateBlock(oldBlock Block, BPM int) Block {
	var newBlock Block
	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.BPM = BPM
	newBlock.Timestamp = t.String()
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateBlockHash(newBlock)

	return newBlock
}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateBlockHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func makeBasicHost(listenPort int, secio bool, randseed int64) (host.Host, error) {
	// if the seed is zero, use real cryptographic randomness. Otherwise, use a deterministic randomness source to make generated keys
	// stay the same accross multiple runs

	var r io.Reader

	if randseed == 0 {
		r = rand.Reader
	} else {
		r = mrand.New(mrand.NewSource(randseed))
	}

	// Generata a key pair for the host. We will use it
	// to obtain a valid host ID.

	priv, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)

	if err != nil {
		return nil, err
	}

	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", listenPort)),
		libp2p.Identity(priv),
	}

	if !secio {
		opts = append(opts, libp2p.NoSecurity)
	}

	basicHost, err := libp2p.New(context.Background(), opts...)
	if err != nil {
		return nil, err
	}

	// Build host multiaddress
	hostAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", basicHost.ID().Pretty()))

	// Now we can build a full multiaddress to reach this host
	// by encapsulating both addresses;

	addr := basicHost.Addrs()[0]
	fullAddr := addr.Encapsulate(hostAddr)

	log.Printf("I am %s\n", fullAddr)

	if secio {
		log.Printf("Now run \" go run main.go -l %d -d %s -secio \"  on a different terminal\n", listenPort+1, fullAddr)
	} else {
		log.Printf("Now run \" go run main.go -l %d -d %s \" on a different terminal\n", listenPort+1, fullAddr)
	}

	return basicHost, nil
}

func handerStream(s net.Stream) {
	log.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))
	go readData(rw)
	go writeData(rw)
}

func readData(rw *bufio.ReadWriter) {
	for {
		str, err := rw.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		if str == "" {
			return
		}

		if str != "\n" {
			chain := make([]Block, 0)
			if err := json.Unmarshal([]byte(str), &chain); err != nil {
				log.Fatal(err)
			}

			mutex.Lock()
			if len(chain) > len(Blockchain) {
				Blockchain = chain
				bytes, err := json.MarshalIndent(Blockchain, "", " ")
				if err != nil {
					log.Fatal(err)
				}

				// Green console color:
				// Reset console color
				fmt.Printf("\x1b[32m%s\x1b[0m> ", string(bytes))
			}

			mutex.Unlock()
		}
	}
}

func writeData(rw *bufio.ReadWriter) {
	go func() {
		for {
			time.Sleep(5 * time.Second)
			mutex.Lock()
			bytes, err := json.Marshal(Blockchain)
			if err != nil {
				log.Println(err)
			}

			mutex.Unlock()

			mutex.Lock()
			rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))
			rw.Flush()
			mutex.Unlock()
		}
	}()

	stdReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		sendData = strings.Replace(sendData, "\n", "", -1)
		bpm, err := strconv.Atoi(sendData)
		if err != nil {
			log.Fatal(err)
		}

		lastBlock := Blockchain[len(Blockchain)-1]
		newBlock := generateBlock(lastBlock, bpm)

		if isBlockValid(newBlock, lastBlock) {
			mutex.Lock()
			Blockchain = append(Blockchain, newBlock)
			mutex.Unlock()
		}

		bytes, err := json.Marshal(Blockchain)
		if err != nil {
			log.Println(err)
		}

		spew.Dump(Blockchain)
		mutex.Lock()
		rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))
		rw.Flush()
		mutex.Unlock()
	}
}

func Run() {
	t := time.Now()

	genesisBlock := Block{}
	genesisBlock = Block{0, t.String(), 0, calculateBlockHash(genesisBlock), ""}

	Blockchain = append(Blockchain, genesisBlock)

	// Libp2p code uses golog to log messages, They log with different
	// string IDs. We can control the verbosity level for all loggers with:

	golog.SetAllLoggers(gologging.INFO) //change to DEBUG for extra info

	// Parse options from the command line
	listenF := flag.Int("l", 0, "wait for incoming connections")
	target := flag.String("d", "", "target peer to dail")
	secio := flag.Bool("secio", false, "enable secio")
	seed := flag.Int64("seed", 0, "set random seed for id generation")
	flag.Parse()

	if *listenF == 0 {
		log.Fatal("Please provide a port to bind on with -l")
	}

	// Make a host that listen on the given multiaddress
	ha, err := makeBasicHost(*listenF, *secio, *seed)
	if err != nil {
		log.Fatal(err)
	}

	if *target == "" {
		log.Println("listening for connections")
		// Set a stream handler on host A. /p2p/1.0.0 is
		// a user-defined protocol name
		ha.SetStreamHandler("/p2p/1.0.0", handerStream)
		select {} //hang forever
	} else {
		ha.SetStreamHandler("/p2p/1.0.0", handerStream)

		// the following code extracts target's peer ID from the
		// given multiaddress

		ipfsaddr, err := ma.NewMultiaddr(*target)
		if err != nil {
			log.Fatal(err)
		}

		pid, err := ipfsaddr.ValueForProtocol(ma.P_IPFS)
		if err != nil {
			log.Fatal(err)
		}

		peerid, err := peer.IDB58Decode(pid)
		if err != nil {
			log.Fatalln(err)
		}

		// Decapsulate the /ipfs/<peerId> part from the target
		// ip4/<a.b.c.d>/ipfs/<peer> becames /ip4/<a.b.c.d>

		targetPeerAddr, _ := ma.NewMultiaddr(
			fmt.Sprintf("/ipfs/%s", peer.IDB58Encode(peerid)))

		targetAddr := ipfsaddr.Decapsulate(targetPeerAddr)

		// We have a peer ID and a a targetAddr so we add it to the peerstore
		// so LibP2P knows how to contact it

		ha.Peerstore().AddAddr(peerid, targetAddr, pstore.PermanentAddrTTL)

		log.Println("opening a stream")
		// make a new stream from the host B to host A
		// it should be handled on host A by the handler we set above bacause
		// we use the same /p2p/1.0.0 protocol
		s, err := ha.NewStream(context.Background(), peerid, "/p2p/1.0.0")
		if err != nil {
			log.Fatalln(err)
		}

		// Create a buffered stream so that read and writes are non blocking.
		rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

		go writeData(rw)
		go readData(rw)

		select {} // hang forever

	}
}
