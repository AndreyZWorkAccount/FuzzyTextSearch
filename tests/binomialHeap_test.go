//Copyright Copyright 2018 Andrey Z.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0

package tests

import (
	"math/rand"
	"testing"

	"github.com/AndreyZWorkAccount/Levenshtein/priorityQueue"
)

type ValueWithPriority struct {
	value, priority uint
}

func (v *ValueWithPriority)	Value() interface{}{
	return v.value
}

func (v *ValueWithPriority) Priority() uint{
	return v.priority
}




func TestInsertIntoBinHeap(t *testing.T) {

	heap := priorityQueue.NewBinomialHeap()

	if heap == nil{
		t.Error("Empty heap should be not nil.")
	}

	if heap.Size() != 0{
		t.Error("Empty heap should have zero size.")
	}

	//Insert single value
	newVal := ValueWithPriority{value:1, priority:1}
	heap.Insert(&newVal)
	//assert
	sizeShouldBe(heap, 1, t)
	shouldHaveTreeWithRank(heap, 0, t)


	//Insert new value
	newVal = ValueWithPriority{value:1, priority:2}
	heap.Insert(&newVal)
	//assert
	sizeShouldBe(heap, 2, t)
	shouldNotHaveTreeWithRank(heap, 0, t)
	shouldHaveTreeWithRank(heap, 1, t)


	//Insert new value
	newVal = ValueWithPriority{value:1, priority:0}
	heap.Insert(&newVal)
	//assert
	sizeShouldBe(heap, 3, t)
	shouldHaveTreeWithRank(heap, 0, t)
	shouldHaveTreeWithRank(heap, 1, t)

	//Insert new value
	newVal = ValueWithPriority{value:1, priority:0}
	heap.Insert(&newVal)
	//assert
	sizeShouldBe(heap, 4, t)
	shouldNotHaveTreeWithRank(heap, 0, t)
	shouldNotHaveTreeWithRank(heap, 1, t)
	shouldHaveTreeWithRank(heap, 2, t)
}

func TestPopFromBinHeap(t *testing.T){
	heap := priorityQueue.NewBinomialHeap()

	var newVal priorityQueue.IPrioritized

	r := rand.New(rand.NewSource(99))

	for i :=uint(100000);i>1;i--{
		//Insert
		newVal = &ValueWithPriority{value:uint(r.Intn(100000)), priority:uint(r.Intn(100000))}
		heap.Insert(newVal)
	}

	current := heap.Peek()

	for heap.Size() > 0{
		tmp := heap.Pop()

		if tmp.Priority() < current.Priority(){
			t.Error("Elemenst order is broken.")
		}
		current = tmp
	}
}

func TestPopDecreaseSize(t *testing.T){
	heap := priorityQueue.NewBinomialHeap()

	//Insert
	newVal := &ValueWithPriority{value:100, priority:100}
	heap.Insert(newVal)

	sizeShouldBe(heap,1, t)

	//Pop
	heap.Pop()

	sizeShouldBe(heap,0, t)
}



