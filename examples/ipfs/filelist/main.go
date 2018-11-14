/**
This is a script to list all files which in ipfs repo
**/

package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path"
	// "sort"
	dshelp "gx/ipfs/QmTmqJGRQfuH8eKWD1FjThwPRipt1QhqJQNZ8MpzmfAAxo/go-ipfs-ds-help"
	ds "gx/ipfs/QmXRKBQA4wXP7xWbFiZsR1GP4HV6wMDQ1aWFxZZ4uBcPX9/go-datastore"
	"strings"
)

const (
	// BaseDir is the path of ipfs repo
	BaseDir = "/Users/magic/.ipfs"
)

func main() {
	p := path.Join(BaseDir, "blocks")
	files, _ := ioutil.ReadDir(p)
	res := IpfsRes()

	for _, f := range files {
		if f.IsDir() {
			sfiles, _ := ioutil.ReadDir(path.Join(p, f.Name()))
			for _, sf := range sfiles {
				s := strings.Split(sf.Name(), ".")
				sfName := s[0]
				k := KeyToCid(sfName)
				for _, c := range res {
					if c == k {
						fmt.Printf("%s 在列表中 \n", k)
					}
				}
			}
		}
	}

}

// KeyToCid 转换keyToCid
func KeyToCid(key string) (cid string) {
	newKey := ds.NewKey(key)
	c, err := dshelp.DsKeyToCid(newKey)
	if err != nil {
		return
	}
	return fmt.Sprintf("%s", c)

}

// IpfsRes 获取ipfs结果
func IpfsRes() []string {

	res := make([]string, 100)
	cmd := exec.Command("/bin/bash", "-c", `ipfs refs local`)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error:", err)
		return res
	}
	if err = cmd.Start(); err != nil {
		fmt.Println("Error", err)
		return res
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("Error:", err)
		return res
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println("Error:", err)
		return res
	}
	for _, s := range strings.Split(string(bytes), "\n") {
		res = append(res, s)
	}
	return res
}
