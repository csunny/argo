/*
This is tool for blocks dir file transfer ipfs cid. 
*/
package main

import (
	"fmt"
	cid "gx/ipfs/QmcZfnkapfECQGcLZaf9B79NRg7cRa9EnZh4LSbkCzwNvY/go-cid"
	ds "gx/ipfs/QmXRKBQA4wXP7xWbFiZsR1GP4HV6wMDQ1aWFxZZ4uBcPX9/go-datastore"
	dshelp  "gx/ipfs/QmTmqJGRQfuH8eKWD1FjThwPRipt1QhqJQNZ8MpzmfAAxo/go-ipfs-ds-help"
)

func main(){
	k := CidToKey("QmcHv3rHetEbsWkmmEJqJstssC2RuufqN8H9GMoK9W6Hop")
	fmt.Println("cid to key", k)

	c := KeyToCid("CIQBPSVUOU3PYL6WUJ552TRS26JLONKVHJAFXR6Z5K6BF5VAWZ55R7A")
	fmt.Println("key to cid", c)
}


// CidToKey 转换cid到key。
func CidToKey(str string) (key string){
	
	c, _ := cid.Decode(str)
	dsKey := dshelp.CidToDsKey(c)
	return fmt.Sprintf("%s", dsKey)
}

// KeyToCid 转换keyToCid
func KeyToCid(key string) (cid string){
	newKey := ds.NewKey(key)
	c, err := dshelp.DsKeyToCid(newKey)
	if err != nil{
		return 
	}
	return fmt.Sprintf("%s", c)
}