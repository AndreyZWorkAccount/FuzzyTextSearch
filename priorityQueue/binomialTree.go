//Copyright Copyright 2018 Andrey Z.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0

package priorityQueue

type Rank int

type BinomialTree struct {

	rank Rank

	root BinomialNode
}

func NewBinomialTree(val IPrioritized) BinomialTree{
	return BinomialTree{
		rank:0,
		root:newBinomialNode(val.Priority(), val.Value()),
	}
}

func (bt *BinomialTree) merge(other *BinomialTree) BinomialTree{
	mergedRoot := bt.root.mergeWith(&other.root)
	return 	BinomialTree{ root: *mergedRoot, rank: bt.rank  + 1}
}