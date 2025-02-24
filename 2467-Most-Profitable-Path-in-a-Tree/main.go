package main

import (
	"fmt"
	"math"
)

func main() {
	edges := [][]int{{0, 2}, {0, 6}, {1, 3}, {1, 5}, {2, 5}, {4, 6}}
	bob := 6
	amount := []int{8838, -6396, -5940, 2694, -1366, 4616, 2966}
	fmt.Println(mostProfitablePath(edges, bob, amount))
}

type MyTreeNode struct {
	Children  []*MyTreeNode
	Parent    *MyTreeNode
	Val       int
	Cost      int
	BobStep   int
	Hierarchy int
}

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	// prepare the nodes
	allNodes := make([]*MyTreeNode, len(amount))
	for i := 0; i < len(amount); i++ {
		node := &MyTreeNode{Val: i, Cost: amount[i]}
		node.BobStep = -1
		allNodes[i] = node
	}

	allNodes[0].Hierarchy = len(allNodes)

	// link nodes by their edges
	graph := make([][]int, len(amount))
	for _, e := range edges {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	// reconstruct the tree
	buildTree(graph, 0, allNodes)

	// now we will calculate Bob path first
	bobsCurrentNode := allNodes[bob]
	for bobStepI := 0; ; bobStepI++ {
		bobsCurrentNode.BobStep = bobStepI
		if bobsCurrentNode.Val == 0 {
			break
		}
		bobsCurrentNode = bobsCurrentNode.Parent
	}

	// finally we can perform dfs to calculate best Alices Path
	root := allNodes[0]
	currentIncome := 0
	bestIncome := math.MinInt32
	aliceStep := 0
	dfs(root, aliceStep, currentIncome, &bestIncome)
	return bestIncome
}

func dfs(tn *MyTreeNode, aliceStep int, currentIncome int, bestIncome *int) {
	type stackItem struct {
		node          *MyTreeNode
		aliceStep     int
		currentIncome int
	}

	stack := []stackItem{
		{node: tn, aliceStep: aliceStep, currentIncome: currentIncome},
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node := top.node
		step := top.aliceStep
		income := top.currentIncome

		newIncome := income
		if step < node.BobStep || node.BobStep == -1 {
			newIncome += node.Cost
		} else if step == node.BobStep {
			newIncome += node.Cost / 2
		}

		if len(node.Children) == 0 {
			if newIncome > *bestIncome {
				*bestIncome = newIncome
			}
			continue
		}

		for _, child := range node.Children {
			stack = append(stack, stackItem{
				node:          child,
				aliceStep:     step + 1,
				currentIncome: newIncome,
			})
		}
	}
}

func buildTree(graph [][]int, root int, nodes []*MyTreeNode) {
	nodes[root].Parent = nil

	visited := make([]bool, len(nodes))
	visited[root] = true

	queue := []int{root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, nxt := range graph[curr] {
			if !visited[nxt] {
				visited[nxt] = true
				nodes[nxt].Parent = nodes[curr]
				nodes[curr].Children = append(nodes[curr].Children, nodes[nxt])
				queue = append(queue, nxt)
			}
		}
	}
}
