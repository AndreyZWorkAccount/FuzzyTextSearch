//Copyright Copyright 2018 Andrey Z.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"fmt"
	"github.com/AndreyZWorkAccount/Levenshtein/levenshteinAlg"
	async "github.com/AndreyZWorkAccount/Levenshtein/levenshteinAsync"
	"time"
)

func main() {
	//setup
	const testWord = "silly"
	const dictionaryFileName = "main\\words"
	const dictionarySize = 256
	const topCount = 20
	requestProcessingTime := time.Second * 2
	costs := levenshteinAlg.ChangesCosts{1,1,1}

	fmt.Printf("Word to search: %v.\n\n", testWord)

	//read input
	ok, dictionaries := readDictionaries(dictionaryFileName, dictionarySize)
	if !ok{
		return
	}

	//run processor
	processor := async.NewProcessor(dictionaries, requestProcessingTime, costs)
	processor.Start()

	//send request
	processor.Requests() <- async.NewRequest(testWord)
	result := waitForResponse(requestProcessingTime, processor.Responses())

	if result == nil{
		return
	}

	fmt.Println("Most matching:")
	for _, res := range result.GetItems(topCount) {
		fmt.Printf("%v  ( distance:  %v ).\n", res.Word, res.Distance)
	}
}

func waitForResponse(requestProcessingTime time.Duration, responses <-chan async.Response) async.Response {
	defer timeTrack(time.Now(),"waitForResponse")

	requestBreak := time.After(requestProcessingTime)
	for {
		select {
		case response := <-responses:
		    return response
		case <-requestBreak:
			fmt.Println("Processing timeout.")
			return nil
		default:
		}
	}
}

func timeTrack(start time.Time, operation string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n\n", operation, elapsed)
}

