//Copyright Copyright 2018 Andrey Z.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0

package tests

import (
	. "github.com/AndreyZWorkAccount/Levenshtein/priorityQueue"
	"testing"
)

func shouldHaveTreeWithRank(bh *BinomialHeap, rank Rank, t *testing.T) {
	if !bh.HasTreeWithRank(rank) {
		t.Errorf("There should be a binomial tree with rank %v.", rank)
	}
}

func shouldNotHaveTreeWithRank(bh *BinomialHeap, rank Rank, t *testing.T) {
	if bh.HasTreeWithRank(rank) {
		t.Errorf("There shouldn't be a binomial tree with rank %v.", rank)
	}
}

func sizeShouldBe(bh *BinomialHeap, size uint, t *testing.T) {
	if bh.Size() != size {
		t.Errorf("Size should be equal to %v.", size)
	}
}
