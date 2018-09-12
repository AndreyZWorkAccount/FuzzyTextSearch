//Copyright Copyright 2018 Andrey Z.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0

package priorityQueue

import "github.com/AndreyZWorkAccount/Levenshtein/extensions"

func (bh *BinomialHeap) getTreeWithRank(rank Rank) *BinomialNode{
	for _,n := range bh.forest{
		if n.rank == rank{
			return n
		}
	}
	return nil
}

func (bh *BinomialHeap) getMinimumTree() *BinomialNode{
	if len(bh.forest) == 0{
		return nil
	}

	minPriority := extensions.MaxUInt
	var ansRank Rank
	for rank, tree := range bh.forest{
		if tree.priority < minPriority{
			minPriority = tree.priority
			ansRank = rank
		}
	}
	res := bh.forest[ansRank]
	return res
}


