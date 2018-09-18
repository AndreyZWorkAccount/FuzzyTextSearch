//Copyright Copyright 2018 Andrey Z.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0

package tests

import (
	"github.com/AndreyZWorkAccount/FuzzyTextSearch/fuzzySearch"
	"github.com/AndreyZWorkAccount/FuzzyTextSearch/levenshteinAlg"
	"sort"
	"testing"
)

func TestResponseBinomialMerge(t *testing.T) {

	respOneDistances := [...]levenshteinAlg.Distance{
		levenshteinAlg.Distance{Value: 1, ToWord: "A"},
		levenshteinAlg.Distance{Value: 7, ToWord: "AAAAAAA"},
	}

	respTwoDistances := [...]levenshteinAlg.Distance{
		levenshteinAlg.Distance{Value: 2, ToWord: "AA"},
		levenshteinAlg.Distance{Value: 5, ToWord: "AAAAA"},
	}

	etalonDistances := append(respOneDistances[:], respTwoDistances[:]...)
	sort.Slice(etalonDistances, func(i, j int) bool {
		return etalonDistances[i].Value < etalonDistances[j].Value
	})

	responceOne := fuzzySearch.NewResponseBinomial(respOneDistances[:])
	responceTwo := fuzzySearch.NewResponseBinomial(respTwoDistances[:])

	responceOne.Merge(responceTwo)

	actualDistances := responceOne.GetItems()

	if len(actualDistances) != len(etalonDistances) {
		t.Error("Incorrect length of merged result.")
	}

	for i := 0; i < len(actualDistances); i++ {
		actual := actualDistances[i]
		etalon := etalonDistances[i]

		if actual.Distance != etalon.Value {
			t.Error("Distances should be equal.")
		}

		if actual.Word != etalon.ToWord {
			t.Error("Words should be equal.")
		}
	}

}
