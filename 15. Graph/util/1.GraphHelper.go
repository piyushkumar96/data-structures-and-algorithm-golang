package util

type Node struct {
	Id    interface{}
	Edges [][2]interface{}
}

type Graph struct {
	Nodes           map[interface{}]*Node
	NodesKeyInOrder []interface{}
}

func CreateGraph() *Graph {
	return &Graph{
		Nodes:           make(map[interface{}]*Node),
		NodesKeyInOrder: make([]interface{}, 0),
	}
}

func (g *Graph) AddNode(id interface{}) {
	g.Nodes[id] = &Node{
		Id:    id,
		Edges: make([][2]interface{}, 0),
	}
	g.NodesKeyInOrder = append(g.NodesKeyInOrder, id)
}

func (g *Graph) AddEdge(n1 interface{}, n2 interface{}, w interface{}) {
	g.Nodes[n1].Edges = append(g.Nodes[n1].Edges, [2]interface{}{n2, w})

}

func (g *Graph) GetNodes() []interface{} {
	nodes := make([]interface{}, 0, 100)

	for _, v := range g.NodesKeyInOrder {
		nodes = append(nodes, v)
	}

	return nodes
}

func (g *Graph) GetEdges() [][3]interface{} {
	edges := make([][3]interface{}, 0, len(g.Nodes))

	for _, v1 := range g.NodesKeyInOrder {
		for _, v2 := range g.Nodes[v1].Edges {
			edges = append(edges, [3]interface{}{v1, v2[0], v2[1]})
		}
	}
	return edges
}

func (g *Graph) GetEdgesForNode(id interface{}) [][3]interface{} {
	edges := make([][3]interface{}, 0, len(g.Nodes[id].Edges))
	for _, v2 := range g.Nodes[id].Edges {
		edges = append(edges, [3]interface{}{id, v2[0], v2[1]})
	}
	return edges
}
