package main

import (
	"fmt"
	"strconv"
)

func main() {
	traversal := "1-5--9---1----4--5---3-6"
	root := recoverFromPreorder(traversal)
	fmt.Println(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverFromPreorder(S string) *TreeNode {
	var rootVal int
	for i, r := range S {
		if r == '-' {
			rootVal, _ = strconv.Atoi(S[:i])
			root := &TreeNode{Val: rootVal}
			dfsWithDashes(root, S[i:]+"-", 1)
			return root
		}
	}
	rootVal, _ = strconv.Atoi(string(S[0]))
	root := &TreeNode{Val: rootVal}
	return root
}

func dfsWithDashes(tn *TreeNode, s string, expectedD int) int {
	realD := 0
	i := 0
	skip := 0
	for ; i < len(s); i++ {
		if s[i] == '-' {
			if realD == expectedD {

				if tn.Left == nil {
					val, err := strconv.Atoi(s[skip+realD : i])
					if err != nil {
						continue
					}
					child := &TreeNode{Val: val}
					tn.Left = child
					skip = dfsWithDashes(child, s[i:], expectedD+1)
					i = i + skip
					realD = 1
					continue
				}
				if tn.Right == nil {
					lChildLen := len(strconv.Itoa(tn.Left.Val))
					val, err := strconv.Atoi(s[skip+realD*2+lChildLen : i])
					if err != nil {
						continue
					}
					child := &TreeNode{Val: val}
					tn.Right = child
					skip = dfsWithDashes(child, s[i:], expectedD+1)
					i = i + skip
					return i
				}
				return i - realD
			} else {
				realD += 1
			}
		} else {
			if realD != expectedD {
				return i - realD
			}
		}
	}

	return len(s)
}
