package graph

import "fmt"

type Node struct {
	Id int
	Label string
}

type Edge struct {
	From int
	To int
	Weight int
}

type Graph struct {
	Nodes []Node
	Edges [][]Edge
}

func createGraph(n int) *Graph {
	return &Graph{
		Nodes: make([]Node, n),
		Edges: make([][]Edge, n),
	}
}

func (g *Graph) addEdge(from, to, weight int) {
	g.Edges[from] = append(g.Edges[from], Edge{From: from, To: to, Weight: weight})
}

func (g *Graph) addNode(id int, label string) {
	g.Nodes[id] = Node{Id: id, Label: label}
}

func (g *Graph) adjacentEdges(nodeId int) {
	if (len(g.Edges[nodeId]) == 0) {
		return
	}
	fmt.Printf("Printing all edges adjacent to node with id %d\n", nodeId)
	for _, edge := range g.Edges[nodeId] {
		fmt.Printf("Edge %d -> %d - weight %d\n", nodeId, edge.To, edge.Weight)
	}
}

func (g *Graph) BFS(sourceId int) {
	visited := map[int]bool{}
	queue := make([]int, 0)

	visited[sourceId] = true
	queue = append(queue, sourceId)

	fmt.Printf("BFS traversal starting from node with id %d\n", sourceId)

	for len(queue) != 0 {
		s := queue[0]
		fmt.Printf(" - %d %s\n", s, g.Nodes[s].Label)
		queue = queue[1:]

		for _, v := range g.Edges[s] {
			if (!visited[v.To]) {
				visited[v.To] = true
				queue = append(queue, v.To)
			}
		}
	}
}

func (g* Graph) DFSHelper(sourceId int, visited map[int]bool) {
	visited[sourceId] = true
	fmt.Printf(" - %d %s\n", sourceId, g.Nodes[sourceId].Label)

	for _, v := range g.Edges[sourceId] {
		if (!visited[v.To]) {
			g.DFSHelper(v.To, visited)
		}
	}
}

func (g *Graph) DFS(sourceId int) {
	visited := map[int]bool{}

	fmt.Printf("DFS traversal starting from node with id %d\n", sourceId)

	g.DFSHelper(sourceId, visited)
}

