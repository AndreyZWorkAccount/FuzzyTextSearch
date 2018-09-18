//Copyright Copyright 2018 Andrey Z.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0

package levenshteinAsync

import (
	"github.com/AndreyZWorkAccount/Levenshtein/priorityQueue"
)

type ResponseBinomial struct {
	heap *priorityQueue.BinomialHeap
}

//Response implementation
func (r *ResponseBinomial) GetItems(count uint) []ResponseItem {
	result := make([]ResponseItem, 0)

	ok, item := r.heap.Pop()
	for ; ok && count > 0; ok, item = r.heap.Pop() {
		newItem := ResponseItem{item.Value().(string), item.Priority()}
		result = append(result, newItem)
		count--
	}

	return result
}

func (r *ResponseBinomial) Merge(other Response) {
	otherBinomial, ok := other.(*ResponseBinomial)
	if !ok {
		panic("Merge is impossible.")
	}
	r.heap.Merge(otherBinomial.heap)
}
