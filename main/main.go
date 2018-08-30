//Copyright Copyright 2018 Andrey Z.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"bufio"
	"strings"
	"fmt"

	"github.com/AndreyZWorkAccount/Levenshtein/trie"
	"github.com/AndreyZWorkAccount/Levenshtein/levenshteinAlg"
)

func main() {
	const dict = "ax ab abc abcd abcde abcdef abcdefg"
	const testWord = "ab"

	//read input
	scanner := bufio.NewScanner(strings.NewReader(dict))
	scanner.Split(bufio.ScanWords)

	//fill vocabulary
	vocabulary := trie.New()
	for scanner.Scan(){
		vocabulary.Put(scanner.Text())
	}

	for _,w := range vocabulary.Words(){
		fmt.Println(w)
	}

	fmt.Println(levenshteinAlg.Run(vocabulary, testWord))
}


