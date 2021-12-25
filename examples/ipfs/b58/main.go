package main

import (
	"fmt"
	"github.com/multiformats/go-multihash"
)

// ID magic test id convert
type ID string


func main(){
	m, err := multihash.FromB58String("QmNsDs3LCDsvGaPazxdNw3izm2bQ1enwMYhEtRpRrPvHyX")
	if err != nil{
		fmt.Println(err)
	}

	res := ID(m)
	fmt.Println(res)
}