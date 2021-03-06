//Copyright Copyright 2018 Andrey Z.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
package levenshteinAlg

import (
	. "github.com/AndreyZWorkAccount/FuzzyTextSearch/extensions"
	"github.com/AndreyZWorkAccount/FuzzyTextSearch/trie"
)

func Run(node trie.INode, word string, costs ChangesCosts) []Distance {
	editDistances := make([]uint, len(word)+1)
	for k := range editDistances {
		editDistances[k] = uint(k)
	}
	distance := make([]Distance, 0)

	for _, c := range node.Children() {
		input := inputArgs{c, []rune(word)}

		var newDistances []Distance

		initialDistances := append([]uint(nil), editDistances...)
		context := stepContext{initialDistances, make([]trie.INode, 0)}

		newDistances = run(input, context, &costs)

		distance = append(distance, newDistances...)
	}
	return distance
}

func run(input inputArgs, context stepContext, costs *ChangesCosts) (outRes []Distance) {
	result := make([]Distance, 0)

	word := input.word
	node := input.node
	previousDistances := context.distances

	currentDistances := calcCurrentDistances(node, word, previousDistances, costs)

	if node.IsFinal() {
		currentWordDistance := currentDistances[len(currentDistances)-1]
		visNodes := append(context.visitedNodes, node)
		result = append(result, Distance{currentWordDistance, GetWord(visNodes)})
	}

	children := node.Children()
	if len(children) == 0 {
		return result
	}

	for _, n := range children {
		input := inputArgs{n, word}
		context := stepContext{currentDistances, append(context.visitedNodes, node)}

		newDistances := run(input, context, costs)
		result = append(result, newDistances...)
	}

	return result
}

func calcCurrentDistances(node trie.INode, word []rune, previousDistances []uint, costs *ChangesCosts) []uint {
	currentDistances := make([]uint, len(word)+1)
	currentDistances[0] = previousDistances[0] + 1
	lettersCount := len(word) + 1

	letter := node.Symbol()

	for pos := 1; pos < lettersCount; pos++ {

		addDist := previousDistances[pos] + costs.AddCost
		removeDist := currentDistances[pos-1] + costs.RemoveCost
		repDist := previousDistances[pos-1]
		if word[pos-1] != letter {
			repDist += costs.ReplaceCost
		}
		currentDistances[pos] = Min(removeDist, addDist, repDist)
	}
	return currentDistances
}

func GetWord(nodes []trie.INode) string {
	runes := make([]rune, 0)
	for _, n := range nodes {
		runes = append(runes, n.Symbol())
	}
	return string(runes)
}
