//Copyright Copyright 2018 Andrey Z.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0

package fuzzySearch

import (
	"sync"
	"time"

	"github.com/AndreyZWorkAccount/FuzzyTextSearch/levenshteinAlg"
)

func (p *RequestsProcessor) processRequest(request SearchRequest) Response {
	response := NewResponse(make([]levenshteinAlg.Distance, 0))

	allDone, responseChan := p.waitForAllDictionaries(request)

	timeout := time.After(p.requestTimeout)
	for {
		select {
		case newResponse := <-responseChan:
			response.Merge(newResponse)
		case <-allDone:
			return response
		case <-timeout:
			return response
		}
	}
}

func (p *RequestsProcessor) waitForAllDictionaries(request SearchRequest) (allDone <-chan struct{}, responses <-chan Response) {
	waitChan := make(chan struct{})
	responseChanel := make(chan Response)

	wg := sync.WaitGroup{}
	wg.Add(len(p.dictionaries))

	for _, dict := range p.dictionaries {
		searchDictionary := dict
		go func() {
			defer wg.Done()
			searchResult := levenshteinAlg.Run(searchDictionary, request.word, p.costs)
			responseChanel <- NewResponse(searchResult)
		}()
	}

	go func() {
		wg.Wait()
		close(waitChan)
	}()

	allDone = waitChan
	responses = responseChanel
	return
}
