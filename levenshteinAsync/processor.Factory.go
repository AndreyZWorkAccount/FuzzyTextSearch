//Copyright Copyright 2018 Andrey Z.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0

package levenshteinAsync

import (
	"github.com/AndreyZWorkAccount/Levenshtein/levenshteinAlg"
	"github.com/AndreyZWorkAccount/Levenshtein/trie"
	"time"
)

//Factory method
func NewProcessor(dictionaries []trie.INode, timeout time.Duration, costs levenshteinAlg.ChangesCosts) IProcessor {

	requests := make(chan SearchRequest)
	responses := make(chan Response)
	cancellation := make(chan struct{})

	return &RequestsProcessor{costs, timeout, dictionaries, requests, responses, cancellation}
}
