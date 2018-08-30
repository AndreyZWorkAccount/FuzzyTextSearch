//Copyright Copyright 2018 Andrey Z.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
package levenshteinAlg

import (
	"github.com/AndreyZWorkAccount/Levenshtein/trie"
	. "github.com/AndreyZWorkAccount/Levenshtein/extensions"
)


func Run(node trie.INode, word string) []Distance {
	editDistances := make([]int, len(word)+1)
	for k := range editDistances {
		editDistances[k] = k
	}
	distance := make([]Distance,0)

	for _,c := range node.Children(){
		newDistances := run(c, c.Symbol(), []rune(word), editDistances)
		distance = append(distance, newDistances...)
	}
	return distance
}

func run(node trie.INode, letter rune, word []rune, previousDistances []int) []Distance {
	result := make([]Distance,0)

	currentDistances := make([]int, len(word)+1)
	currentDistances[0] = previousDistances[0] + 1
	lettersCount := len(word) + 1

	for pos := 1; pos < lettersCount; pos++ {

		addDist := previousDistances[pos] + 1
		removeDist := currentDistances[pos-1] + 1
		repDist := previousDistances[pos-1]
		if word[pos-1] != letter {
			repDist += 1
		}
		currentDistances[pos] = Min([]int{removeDist, addDist, repDist})

	}

	if node.IsFinal(){
		currentWordDistance := currentDistances[len(currentDistances)-1]
		result = append(result,Distance{currentWordDistance, string(node.Symbol())})
	}

	children := node.Children()
	if len(children) == 0{
		return result
	}

	for _,n := range children{
		newDistances := run(n, n.Symbol(), word, currentDistances)
		result = append(result, newDistances...)
	}

	return result
}


