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
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/AndreyZWorkAccount/FuzzyTextSearch/trie"
	"github.com/AndreyZWorkAccount/FuzzyTextSearch/vocabularyReader"
)

func readDictionaries(dictionaryFileName string, dictionarySize int) (ok bool, dictionaries []trie.INode) {
	currentDir, _ := filepath.Abs("./")

	b, err := ioutil.ReadFile(currentDir + "\\" + dictionaryFileName) // just pass the file name
	if err != nil {
		fmt.Print(err)
		ok = false
		dictionaries = nil
		return
	}

	scanner := bufio.NewScanner(bytes.NewReader(b))
	scanner.Split(bufio.ScanWords)

	chunks := createInputChunks(scanner, dictionarySize)
	readers := createVocReaders(chunks)

	dictionaries = createVocabularies(readers)
	ok = true
	return
}

func createInputChunks(scanner *bufio.Scanner, chunkSize int) [][]string {
	result := make([][]string, 0)

	canRead := scanner.Scan()
	for canRead {
		chunk := make([]string, 0)
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

func createVocabularies(readers []vocabularyReader.IVocabularyReader) []trie.INode {
	tries := make([]trie.INode, 0)
	for _, reader := range readers {
		trie := trie.New()
		item := reader.ReadElement()
		for item.HasValue {
			trie.Put(item.Value)
			item = reader.ReadElement()
		}
		tries = append(tries, trie)
	}
	return tries
}
