package main

import (
	"errors"
)

type (
	ConnectionTree struct {
		nodes []int
		treeWeight []int
		maxWeight int
	}
)

func InitConnectionTree(usersCount int) ConnectionTree {
	users := make([]int, usersCount)
	for i:=0; i<usersCount; i++ {
		users[i] = i
	}
	return ConnectionTree{
		nodes:      users,
		treeWeight: make([]int, usersCount),
		maxWeight: 0,
	}
}

func (ctr *ConnectionTree) Union(left int, right int) (bool, error) {
	if len(ctr.nodes) < left || len(ctr.nodes) < right {
		return false, errors.New("out of range")
	}
	leftRoot := ctr.root(left)
	rightRoot := ctr.root(right)
	if leftRoot == rightRoot {
		return false, nil
	}

	if ctr.treeWeight[leftRoot] < ctr.treeWeight[rightRoot] {
		ctr.joinTrees(rightRoot, leftRoot)
	} else {
		ctr.joinTrees(leftRoot, rightRoot)
	}
	return true, nil
}

func (ctr ConnectionTree) IsInterconnected() bool {
	return ctr.maxWeight + 1 == len(ctr.nodes)
}

func (ctr *ConnectionTree) joinTrees(big int, small int) {
	ctr.nodes[small] = ctr.nodes[big]
	ctr.treeWeight[big] += ctr.treeWeight[small] + 1
	if ctr.treeWeight[big] > ctr.maxWeight {
		ctr.maxWeight = ctr.treeWeight[big]
	}
}

func (ctr *ConnectionTree) root(item int) int {
	for item != ctr.nodes[item] {
		itemWeight := ctr.treeWeight[item] + 1
		ctr.treeWeight[ctr.nodes[item]] -= itemWeight
		ctr.treeWeight[ctr.nodes[ctr.nodes[item]]] += itemWeight
		ctr.nodes[item] = ctr.nodes[ctr.nodes[item]]
		item = ctr.nodes[item]
	}
	return item
}