package dpos

import (
	"io/ioutil"
	"log"
	"strings"
	"strconv"
	"sort"
)

const BPCount = 5

func PickWinner() (bp string) {
	// 选择BlockProducer

	f, err := ioutil.ReadFile(FileName)
	if err != nil {
		log.Fatal(err)
	}

	res := strings.Split(string(f), "\n")

	voteList := make([]int, len(res))
	voteMap := make(map[string]int)
	for _, node := range res {
		nodeSplit := strings.Split(node, ":")
		if len(nodeSplit) > 1 {
			vote, err := strconv.Atoi(nodeSplit[1])
			if err != nil {
				log.Fatal(err)
			}
			voteList = append(voteList, vote)
			voteMap[nodeSplit[0]] = vote
		}
	}
	sort.Slice(voteList, func(i, j int) bool {
		return voteList[i] > voteList[j]
	})

	if len(voteList) > BPCount {
		voteList = voteList[0:BPCount] // 选择前面的5个节点作为Block producer
	}

	for k, v := range voteMap {
		if v > voteList[len(voteList)-1] {
			bp = k
		}
	}
	return
}
