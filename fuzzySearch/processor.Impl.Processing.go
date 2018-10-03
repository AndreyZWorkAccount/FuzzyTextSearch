//Copyright Copyright 2018 Andrey Z.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0

package fuzzySearch

import (
	"github.com/AndreyZWorkAccount/FuzzyTextSearch/trie"
	"time"

	"github.com/AndreyZWorkAccount/FuzzyTextSearch/levenshteinAlg"
)

func (p *RequestsProcessor) processRequest(request SearchRequest) Response {
	response := NewResponse(make([]levenshteinAlg.Distance, 0))

	responseChan := waitForAllDictionaries(request, p.dictionaries, p.costs)

	timeout := time.After(p.requestTimeout)
	expectedRespCnt := len(p.dictionaries)
	receivedRespCnt := 0
	for {
		select {
		case newResponse := <-responseChan:
			receivedRespCnt++
			response.Merge(newResponse)
		    if receivedRespCnt == expectedRespCnt{
		    	return response
			}
		case <-timeout:
			return response
		}
	}
}

func waitForAllDictionaries(request SearchRequest, dictionaries []trie.INode, costs levenshteinAlg.ChangesCosts) (responses <-chan Response) {
	responseChanel := make(chan Response)

	for _, dict := range dictionaries {
		searchDictionary := dict
		go func() {
			searchResult := levenshteinAlg.Run(searchDictionary, request.word, costs)
			responseChanel <- NewResponse(searchResult)
		}()
	}

	responses = responseChanel
	return
}
