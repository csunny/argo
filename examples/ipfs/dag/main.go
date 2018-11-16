/*
* This is the code analysis and use for ipfs dag
*
 */

package main

import (
	"context"
	"fmt"
	offline "gx/ipfs/QmWM5HhdG5ZQNyHQ5XhMdGmV9CvLpFynQfGpTxN2MEM7Lc/go-ipfs-exchange-offline"
	ds "gx/ipfs/QmXRKBQA4wXP7xWbFiZsR1GP4HV6wMDQ1aWFxZZ4uBcPX9/go-datastore"
	bstore "gx/ipfs/QmaG4DZ4JaqEfvPWt5nPPgoTzhc1tr1T3f4Nu9Jpdm8ymY/go-ipfs-blockstore"
	ipld "gx/ipfs/Qme5bWv7wtjUNGsK2BNGVUFPKiuxWrsqrtvYwCLRw8YFES/go-ipld-format"

	bserv "github.com/ipfs/go-ipfs/blockservice"
	dag "github.com/ipfs/go-ipfs/merkledag"
)

func main() {
	ctx := context.Context(context.Background())
	ds := getDagserv()

	data := []byte("This is a test")
	newNode := dag.NodeWithData(data)
	fmt.Println(newNode)

	ds.Add(ctx, newNode)

	node, err := ds.Get(ctx, newNode.Cid())
	if err != nil{
		panic(err)
	}

	// 检验add/get
	fmt.Println(node.Cid().Equals(newNode.Cid()))
}

// getDagServer 创建一个DagServer
func getDagserv() ipld.DAGService {
	dbmap := ds.NewMapDatastore() // 创建一个map 作为datastore
	db := ds.NewLogDatastore(dbmap, "magic")
	
	disk, err := db.DiskUsage()
	if err != nil{
		panic(err)
	}
	fmt.Println("是否使用磁盘持久化  1: 是  0: 否   ===>", disk)

	bs := bstore.NewBlockstore(db) // 新建一个blockstore
	blockserv := bserv.New(bs, offline.Exchange(bs))
	return dag.NewDAGService(blockserv)
}
