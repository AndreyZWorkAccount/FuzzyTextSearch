//Copyright Copyright 2018 Andrey Z.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0

package priorityQueue

type Rank uint

type BinomialNode struct {
	value interface{}

	priority uint

	rank Rank

	children []*BinomialNode
}

func newBinomialNode(priority uint, value interface{}) BinomialNode {
	return BinomialNode{
		priority: priority,
		value:    value,
		rank:     Rank(0),
	}
}

//IPrioritized implementation
func (n *BinomialNode) Priority() uint {
	return n.priority
}

func (n *BinomialNode) Value() interface{} {
	return n.value
}
