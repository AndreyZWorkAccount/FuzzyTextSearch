//Copyright Copyright 2018 Andrey Z.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0

package levenshteinAsync

import (
	"time"

	"github.com/AndreyZWorkAccount/Levenshtein/levenshteinAlg"
	"github.com/AndreyZWorkAccount/Levenshtein/trie"
)

type RequestsProcessor struct {
	costs levenshteinAlg.ChangesCosts
	requestTimeout time.Duration
	dictionaries   []trie.INode
	requests       chan SearchRequest
	responses      chan Response
	cancellation   chan struct{}
}

//IProcessor impl

func (p *RequestsProcessor) Requests() chan SearchRequest {
	return p.requests
}

func (p *RequestsProcessor) Responses() <-chan Response {
	return p.responses
}

func (p *RequestsProcessor) Stop() <-chan bool {
	ch := make(chan bool)
	go p.stopCore(ch)
	return ch
}

func (p *RequestsProcessor) Start() {
	p.startCore()
}

//private
func (p *RequestsProcessor) stopCore(cancellation chan bool) {
	p.cancellation <- struct{}{}
	<-p.cancellation
	cancellation <- true
}

func (p *RequestsProcessor) startCore() {
	cancellation := make(chan struct{})
	go p.runProcessorCore(cancellation)
	p.cancellation = cancellation
}

func (p *RequestsProcessor) runProcessorCore(cancellation chan struct{}) {
	for {
		select {

		case request := <-p.requests:
			p.responses <- p.processRequest(request)
			continue

		case <-cancellation:
			close(cancellation)
			return

		}
	}
}
