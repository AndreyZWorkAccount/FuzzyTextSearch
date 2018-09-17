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
	"fmt"
	"github.com/AndreyZWorkAccount/Levenshtein/vocabularyReader"
	"strings"

	"github.com/AndreyZWorkAccount/Levenshtein/levenshteinAlg"
	"github.com/AndreyZWorkAccount/Levenshtein/trie"
)

func main() {
	const dict = "ax ab abc abcd abcde abcdef abcdefg"
	const testWord = "abce"

	//read input
	scanner := bufio.NewScanner(strings.NewReader(dict))
	scanner.Split(bufio.ScanWords)

	/*chunks := createInputChunks(scanner, 3)
	readers := createVocReaders(chunks)
	vocabularies := createVocabularies(readers)*/

	//fill vocabulary
	vocabulary := trie.New()
	for scanner.Scan() {
		vocabulary.Put(scanner.Text())
	}

	//print vocabulary
	for _, w := range vocabulary.Words() {
		fmt.Println(w)
	}

	//find distances to all words
	fmt.Println(levenshteinAlg.Run(vocabulary, testWord))
}

func createInputChunks(scanner *bufio.Scanner, chunkSize int) [][]string {
	result := make([][]string, 0)

	for scanner.Scan() {
		chunk := make([]string, 0)

		canRead := true
		for canRead && len(chunk) < chunkSize {
			chunk = append(chunk, scanner.Text())
			canRead = scanner.Scan()
		}

		result = append(result, chunk)
	}

	return result
}

func createVocReaders(chunks [][]string) []vocabularyReader.IVocabularyReader {
	readers := make([]vocabularyReader.IVocabularyReader, 0)
	for _, chunk := range chunks {
		reader := vocabularyReader.NewVocReaderStringBased(chunk)
		readers = append(readers, &reader)
	}
	return readers
}

func createVocabularies(readers []vocabularyReader.IVocabularyReader) []*trie.INode {
	tries := make([]*trie.INode, 0)

	for _, reader := range readers {
		trie := trie.New()
		for item := reader.ReadElement(); item.HasValue; {
			trie.Put(item.Value)
		}
	}

	return tries
}
