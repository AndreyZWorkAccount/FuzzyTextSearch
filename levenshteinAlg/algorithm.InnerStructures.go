//Copyright Copyright 2018 Andrey Z.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
package levenshteinAlg

import "github.com/AndreyZWorkAccount/Levenshtein/trie"

type inputArgs struct {
	node trie.INode
	word []rune
}

type stepContext struct {
	distances    []int
	visitedNodes []trie.INode
}
