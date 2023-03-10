package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

type Graph struct {
	nodes []string
	edges []edge
}

type edge struct {
	from, to string
}

func NewGraphFromDependencyMap(m DependencyMap) Graph {
	nodeMap := make(map[string]struct{})
	for k, v := range m {
		nodeMap[k] = struct{}{}

		for _, i := range v {
			nodeMap[i] = struct{}{}
		}
	}

	var nodeList []string
	for i, _ := range nodeMap {
		nodeList = append(nodeList, i)
	}

	edgeMap := make(map[string]struct{})
	for k, v := range m {
		for _, i := range v {
			edgeMap[fmt.Sprintf("%s->%s", k, i)] = struct{}{}
		}
	}

	var edgeList []edge
	for i, _ := range edgeMap {
		split := strings.Split(i, "->")
		edgeList = append(edgeList, edge{
			from: split[0],
			to:   split[1],
		})
	}

	return Graph{
		nodes: nodeList,
		edges: edgeList,
	}
}

func (g Graph) Visualize() error {
	graph := graph.New(graph.StringHash, graph.Directed(), graph.Acyclic())

	for _, node := range g.nodes {
		if err := graph.AddVertex(node); err != nil {
			return fmt.Errorf("couldn't add node to graph: %w", err)
		}
	}

	for _, edge := range g.edges {
		if err := graph.AddEdge(edge.from, edge.to); err != nil {
			return fmt.Errorf("couldn't add edge to graph: %w", err)
		}
	}

	file, err := os.Create("./data.gv")
	if err != nil {
		return fmt.Errorf("couldn't create file for graph: %w", err)
	}

	if err = draw.DOT(graph, file); err != nil {
		return fmt.Errorf("couldn't draw DOT to file: %w", err)
	}

	return nil
}
