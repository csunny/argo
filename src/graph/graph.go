package graph

import (
	"fmt"
)

//  1. 邻接矩阵表达图的数据结构
//  2. 邻接表（矩阵比较稀疏的时候）  V(顶点)  E(边)

// 图的遍历 1. 广度优先  2. 深度优先

// 广度优先遍历： 从图的某一节点出发，首先依次访问该节点的所有相邻顶点，再按照这些顶点被访问的先后次序，依次访问与它们相邻的所有未被访问的顶点。
// 重复此过程，直到所有的顶点均被访问


// 深度优先遍历
// 从图的某个顶点出发访问遍图中所有顶点，且每个顶点仅被访问依次
// DFS
// 1. 访问指定的起始顶点
// 2. 若当前访问的顶点的临接顶点有未被访问的，则任选一个访问，反之，退回到最近访问过的顶点，直到与起始顶点想通的全部顶点都访问完毕
// 3. 若此时图中尚有顶点未被访问， 则选择其中一个顶点作为起始顶点并访问

type Item interface {

}

// 组成图的顶点
type Node struct {
	Value Item
}

// 定义一个图的结构, 图有顶点与边组成 V  E
type ItemGraph struct {
	Nodes []*Node
	Edges map[Node][]*Node
}

func (n *Node) String() string  {
	return fmt.Sprintf("%v", n.Value)
}

// 添加节点
func (g *ItemGraph) AddNode(n *Node)  {
	g.Nodes = append(g.Nodes, n)
}

// 添加边
func (g *ItemGraph) AddEdge(n1, n2 *Node)  {
	if g.Edges == nil{
		g.Edges = make(map[Node][]*Node)
	}

	// 无向图
	g.Edges[*n1] = append(g.Edges[*n1], n2)    // 设定从节点n1到n2的边
	g.Edges[*n2] = append(g.Edges[*n2], n1)    // 设定从节点n2到n1的边
}

func (g *ItemGraph) String()  {
	s := ""

	for i := 0; i< len(g.Nodes); i++{
		s += g.Nodes[i].String() + "->"
		near := g.Edges[*g.Nodes[i]]

		for j :=0; j<len(near); j++{
			s += near[j].String() + " "
		}
		s += "\n"
	}

	fmt.Println(s)
}


// 图的遍历，深度优先与广度优先遍历
// 首先bfs 广度优先搜索
// 此处结合队列实现图的广度优先遍历
type NodeQueue struct {
	Items []Node
}

func (s *NodeQueue) New() *NodeQueue  {
	s.Items = []Node{}
	return s
}

func (s *NodeQueue) Enqueue(t Node)  {
	s.Items = append(s.Items, t)
}

func (s *NodeQueue) Dequeue() *Node {
	item := s.Items[0]
	s.Items = s.Items[1:len(s.Items)]
	return &item
}

func (s *NodeQueue) IsEmpty() bool  {
	return len(s.Items) == 0
}

func (g *ItemGraph) Bfs(f func(node *Node))  {
	q := NodeQueue{}
	q.New()

	n := g.Nodes[0]
	q.Enqueue(*n)

	visited := make(map[*Node]bool)
	visited[n] = true

	for {
		if q.IsEmpty(){
			break
		}
		node := q.Dequeue()
		near := g.Edges[*node]

		for i :=0; i<len(near); i++{
			j := near[i]
			if !visited[j]{
				q.Enqueue(*j)
				visited[j] = true

			}
		}

		if f!=nil{
			f(node)
		}
	}
}

// 以上即为用栈实现图的广度优先遍历



// 下面实现图的深度优先遍历
// ****
//
// implement a stack use go
//
//
// * **//

type NodeStack struct {
	Items []Node
}

func (n *NodeStack) New() *NodeStack {
	//
	n.Items = []Node{}
	return n
}

func (n *NodeStack) push(q Node) {
	n.Items = append(n.Items, q)
}

func (n *NodeStack) pop() *Node {

	item := n.Items[len(n.Items)-1]

	n.Items = n.Items[0: len(n.Items)-1]
	return &item
}

func (n *NodeStack) IsEmpty() bool {
	return len(n.Items) == 0
}

func (n *NodeStack) Size() int {
	return len(n.Items)
}

// DFS implement
func (g *ItemGraph) Dfs(f func(node *Node)) {
	stack := NodeStack{}
	stack.New()

	n := g.Nodes[0]
	stack.push(*n)

	visited := make(map[*Node] bool)

	visited[n] = true

	for {
		if stack.IsEmpty(){
			break
		}

		node := stack.pop()
		if !visited[node]{
			visited[node] = true
		}
		near := g.Edges[*node]

		for i:= 0; i< len(near); i++{
			j := near[i]
			if !visited[j]{
				visited[j] = true
				stack.push(*j)
			}
		}

		if f != nil{
			f(node)
		}
	}
}