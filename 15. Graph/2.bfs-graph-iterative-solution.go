package main

import (
	"./util"
	"fmt"
)

var graph *util.Graph

func bfsIterative(s interface{}) []interface{} {
	resBFS := make([]interface{}, 0, len(graph.Nodes))
	visited := make(map[interface{}]bool, len(graph.Nodes))
	visited[s] = true

	q := util.CreateQueue()

	q.Push(s.(string))

	for !q.IsEmpty() {
		front, _ := q.Pop()
		resBFS = append(resBFS, front)
		for _, v := range graph.GetEdgesForNode(front) {
			if !visited[v[1]] {
				visited[v[1]] = true
				q.Push(v[1].(string))
			}
		}
	}
	return resBFS
}

func main() {
	graph = util.CreateGraph()
	graph.AddNode("S")
	graph.AddNode("A")
	graph.AddNode("B")
	graph.AddNode("C")
	graph.AddNode("D")
	graph.AddNode("E")
	graph.AddNode("F")
	graph.AddNode("G")
	graph.AddEdge("S", "A", 10)
	graph.AddEdge("S", "B", 5)
	graph.AddEdge("S", "C", 5)
	graph.AddEdge("A", "D", 5)
	graph.AddEdge("B", "E", 5)
	graph.AddEdge("C", "F", 5)
	graph.AddEdge("D", "G", 5)
	graph.AddEdge("E", "G", 5)
	graph.AddEdge("F", "G", 5)

	fmt.Print("The graph nodes:- ")
	nodes := graph.GetNodes()
	for _, v := range nodes {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Println("The graph edges")
	edges := graph.GetEdges()
	for _, v := range edges {
		fmt.Println(v)
	}

	fmt.Println("\nBFS Traversal")
	result := bfsIterative("S")
	for _, v := range result {
		fmt.Print(v, " ")
	}

	fmt.Println()
}
