// This is the p2p network, handler the conn and communicate with nodes each other.
// this file is created by magic at 2018-9-2

package dpos

import (
	"io"
	"fmt"
	"os"
	"strings"
	"strconv"
	"time"
	"crypto/rand"
	"flag"
	"log"
	"bufio"
	"sync"
	"encoding/json"
	"context"
	mrand "math/rand"
	"github.com/libp2p/go-libp2p-net"
	"github.com/libp2p/go-libp2p-host"
	"github.com/libp2p/go-libp2p-crypto"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-peer"
	ma "github.com/multiformats/go-multiaddr"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	"github.com/davecgh/go-spew/spew"
)

const DefaultVote = 10
const FileName = "config.ini"

var mutex = &sync.Mutex{}

type Validator struct {
	name string
	vote int
}

func MakeBasicHost(listenPort int, secio bool, randseed int64) (host.Host, error) {
	var r io.Reader

	if randseed == 0 {
		r = rand.Reader
	} else {
		r = mrand.New(mrand.NewSource(randseed))
	}

	// 生产一对公私钥
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

	log.Printf("我是: %s\n", fullAddr)
	SavePeer(basicHost.ID().Pretty())

	if secio {
		log.Printf("现在在一个新终端运行命令: 'go run main/dpos.go -l %d -d %s -secio'\n ", listenPort+1, fullAddr)
	} else {
		log.Printf("现在在一个新的终端运行命令: 'go run main/dpos.go -l %d -d %s '", listenPort+1, fullAddr)
	}
	return basicHost, nil
}

func HandleStream(s net.Stream) {
	log.Println("得到一个新的连接!", s.Conn().RemotePeer().Pretty())
	// 将连接加入到
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
			if len(chain) > len(BlockChain) {
				BlockChain = chain
				bytes, err := json.MarshalIndent(BlockChain, "", " ")
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("\x1b[32m%s\x1b[0m> ", string(bytes))
			}
			mutex.Unlock()
		}
	}
}

func writeData(rw *bufio.ReadWriter) {
	// 启动一个协程处理终端同步
	go func() {
		for {
			time.Sleep(2 * time.Second)
			mutex.Lock()
			bytes, err := json.Marshal(BlockChain)
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

		// pick选择block生产者
		address := PickWinner()
		log.Printf("******节点%s获得了记账权利******", address)
		lastBlock := BlockChain[len(BlockChain)-1]
		newBlock, err := GenerateBlock(lastBlock, bpm, address)
		if err != nil{
			log.Fatal(err)
		}

		if IsBlockValid(newBlock, lastBlock){
			mutex.Lock()
			BlockChain = append(BlockChain, newBlock)
			mutex.Unlock()
		}

		spew.Dump(BlockChain)

		bytes, err := json.Marshal(BlockChain)
		if err != nil {
			log.Println(err)
		}
		mutex.Lock()
		rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))
		rw.Flush()
		mutex.Unlock()
	}
}

func Run() {

	t := time.Now()
	genesisBlock := Block{}
	genesisBlock = Block{0, t.String(), 0, CaculateBlockHash(genesisBlock), "", ""}
	BlockChain = append(BlockChain, genesisBlock)

	// 命令行传参
	listenF := flag.Int("l", 0, "等待节点加入")
	target := flag.String("d", "", "连接目标节点")
	secio := flag.Bool("secio", false, "打开secio")
	seed := flag.Int64("seed", 0, "生产随机数")
	flag.Parse()

	if *listenF == 0 {
		log.Fatal("请提供一个端口号")
	}

	// 构造一个host 监听地址
	ha, err := MakeBasicHost(*listenF, *secio, *seed)
	if err != nil {
		log.Fatal(err)
	}

	if *target == "" {
		log.Println("等待节点连接...")
		ha.SetStreamHandler("/p2p/1.0.0", HandleStream)
		select {}
	} else {
		ha.SetStreamHandler("/p2p/1.0.0", HandleStream)
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
			log.Fatal(err)
		}

		targetPeerAddr, _ := ma.NewMultiaddr(
			fmt.Sprintf("/ipfs/%s", peer.IDB58Encode(peerid)))
		targetAddr := ipfsaddr.Decapsulate(targetPeerAddr)

		// 现在我们有一个peerID和一个targetaddr，所以我们添加它到peerstore中。 让libP2P知道如何连接到它。
		ha.Peerstore().AddAddr(peerid, targetAddr, pstore.PermanentAddrTTL)
		log.Println("打开Stream")

		// 构建一个新的stream从hostB到hostA
		// 使用了相同的/p2p/1.0.0 协议
		s, err := ha.NewStream(context.Background(), peerid, "/p2p/1.0.0")
		if err != nil {
			log.Fatal(err)
		}

		rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

		go writeData(rw)
		go readData(rw)
		select {}
	}
}

func SavePeer(name string) {
	vote := DefaultVote // 默认的投票数目
	f, err := os.OpenFile(FileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()


	f.WriteString(name + ":" + strconv.Itoa(vote) + "\n")

}
