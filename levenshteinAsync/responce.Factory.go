//Copyright Copyright 2018 Andrey Z.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0

package levenshteinAsync

import (
	"github.com/AndreyZWorkAccount/FuzzyTextSearch/levenshteinAlg"
	"github.com/AndreyZWorkAccount/FuzzyTextSearch/priorityQueue"
)

func NewResponse(distances []levenshteinAlg.Distance) *ResponseBinomial {
	heap := priorityQueue.NewBinomialHeap()

	for _, dist := range distances {
		item := priorityQueue.NewPrioritized(dist.ToWord, dist.Value)
		heap.Insert(item)
	}

	return &ResponseBinomial{heap}
}
