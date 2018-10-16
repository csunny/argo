package decision

import (
	"gx/ipfs/QmZUbTDJ39JpvtFCSubiWeUTQRvMA1tVE5RZCJrY4oeAsC/go-ipfs-pq"
	"sync"
	"time"
	wantlist "github.com/csunny/argo/src/ipfs/bitswap/wantlist"
	cid "github.com/ipfs/go-cid"
	peer "github.com/libp2p/go-libp2p-peer"
)

type peerRequestQueue interface{
	Pop() *peerRequestTask
	Push(entry *wantlist.Entry, to peer.ID)
	Remove(k *cid.Cid, p peer.ID)

	// NB
}

func newPRQ() *prq{
	return &prq{
		taskMap: make(map[string]*peerRequestTask),
		partners: make(map[peer.ID]*activePartner),
		frozen: make(map[peer.ID]*activePartner),
		pQueue: pq.New(partnerCompare),
	}
}

var _ peerRequestQueue = &prq{}

type prq struct{
	lock sync.Mutex
	pQueue pq.PQ
	taskMap map[string]*peerRequestTask
	partners map[peer.ID]*activePartner

	frozen map[peer.ID]*activePartner
}

func (tl *prq) Push(entry *wantlist.Entry, to peer.ID){
	tl.lock.Lock()
	defer tl.lock.Unlock()

	partner, ok := tl.partners[to]
	if !ok {
		partner = n
	}
}

type peerRequestTask struct{
	Entry *wantlist.Entry
	Target peer.ID	
	Done func()

	trash bool
	// created marks the time that the task was added to the queue
	created time.Time
	index int  
}
// Key uniquely identifies a task
func (t *peerRequestTask) Key() string{
	return taskKey(t.Target, t.Entry.Cid)
}

func (t *peerRequestTask) Index() int{
	return t.index
}

func (t *peerRequestTask) SetIndex(i int){
	t.index = i
} 

func taskKey(p peer.ID, k *cid.Cid) string{
	return string(p) + k.KeyString()
}

type activePartner struct{
	// active is the number of blocks this peer is currently being sent 
	// active must be locked around as it will be updated externally
	
	actively sync.Mutex
	active int

	activeBlocks *cid.Set
	
	requests int
	index int
	freezeVal int
	taskQueue pq.PQ
}

func newActivePartner() *activePartner{
	return &activePartner{
		taskQueue: pq.New(wrapCmp(V1)),
		activeBlocks: cid.NewSet(),
	}
}