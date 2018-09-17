//Copyright Copyright 2018 Andrey Z.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0

package priorityQueue

//IPriorityQueue implementation
func (b *BinomialHeap) Insert(item IPrioritized) {
	newTree := newBinomialNode(item.Priority(), item.Value())
	b.insert(&newTree)
	b.size++
}

func (b *BinomialHeap) Pop() (ok bool, item IPrioritized) {
	if len(b.forest) == 0 {
		return false, nil
	}
	//find the minimum
	minTree := b.getMinimumTree()
	//remove minimum from heap
	delete(b.forest, minTree.rank)
	//add all child
	for _, c := range minTree.children {
		b.insert(c)
	}

	b.size--

	return true, minTree
}

func (b *BinomialHeap) Peek() IPrioritized {
	if len(b.forest) == 0 {
		return nil
	}
	//find the minimum
	return b.getMinimumTree()
}

func (b *BinomialHeap) Size() uint {
	return b.size
}

func (b *BinomialHeap) Merge(other *BinomialHeap) {
	for _, node := range other.forest {
		b.insert(node)
	}
	b.size += other.size
}
