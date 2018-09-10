//Copyright Copyright 2018 Andrey Z.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0

package priorityQueue

//IPriorityQueue implementation
func (b *BinomialHeap) Insert(item IPrioritized){
	newTree := NewBinomialTree(item)
	b.insert(newTree)
	b.size++
}

func (b *BinomialHeap) Pop() IPrioritized{

	panic("Not implemented")
}

func (b *BinomialHeap) Peek() IPrioritized{

	panic("Not implemented")
}

func (b *BinomialHeap) Size() uint{
	return b.size
}

func (b *BinomialHeap) Merge( other *BinomialHeap) *BinomialHeap{

	panic("Not implemented")
}
