package main

import (
	"fmt"
	mh "gx/ipfs/QmZyZDi491cCNTLfAhwcaDii2Kg4pwKRkhqQzURGDvY6ua/go-multihash"
)

// ID magic test id convert
type ID string


func main(){
	m, err := mh.FromB58String("QmNsDs3LCDsvGaPazxdNw3izm2bQ1enwMYhEtRpRrPvHyX")
	if err != nil{
		fmt.Println(err)
	}

	res := ID(m)
	fmt.Println(res)
}