//Copyright Copyright 2018 Andrey Z.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0

package priorityQueue

func (bh *BinomialHeap) getTreeWithRank(rank Rank) *BinomialTree{
	for _,n := range bh.forest{
		if n.rank == rank{
			return &n
		}
	}
	return nil
}

