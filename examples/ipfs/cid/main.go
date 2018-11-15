/*
* cid is the ipfs content identify module
*  in this module we known how a content convert a hashed id
 */

package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ipfs/go-cid"
)

const (
	// File is the test file
	File = "./argo/docs/readme.md"
)

func main() {

	p := cid.Prefix{
		Version:  1,
		Codec:    0x70,  // prtobuf
		MhType:   0x12,  // sha2-256
		MhLength: -1,
	}

	data, err := ioutil.ReadFile(File)
	if err != nil {
		panic(err)
	}

	fcid, err := p.Sum(data)
	if err != nil{
		panic(err)
	}

	fmt.Println(fcid)

}
