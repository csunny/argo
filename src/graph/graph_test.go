package graph

import (
"testing"
"fmt"
)

var g ItemGraph

func fillGraph()  {
	nA := Node{"A"}
	nB := Node{"B"}
	nC := Node{"C"}
	nD := Node{"D"}
	nE := Node{"E"}
	nF := Node{"F"}

	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddNode(&nE)
	g.AddNode(&nF)

	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nC)
	g.AddEdge(&nB, &nE)
	g.AddEdge(&nC, &nE)
	g.AddEdge(&nE, &nF)
	g.AddEdge(&nD, &nA)
}

func TestAdd(t *testing.T)  {
	fillGraph()
	//g.String()
}

func TestItemGraph_Traverse(t *testing.T) {
	g.Bfs(func(node *Node) {
		fmt.Printf("B-F-S visiting... %v\n", node)
	})
}

func TestItemGraph_Dfs(t *testing.T) {
	g.Dfs(func(node *Node) {
		fmt.Printf("DFS visiting... %v\n", node)
	})
}