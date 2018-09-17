//Copyright Copyright 2018 Andrey Z.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0

package vocabularyReader

type VocReaderStringBased struct {
	sourceStrings []string
	currentIndex  int
}

func NewVocReaderStringBased(strings []string) VocReaderStringBased {
	return VocReaderStringBased{sourceStrings: strings, currentIndex: 0}
}

//IVocabularyReader impl
func (v *VocReaderStringBased) ReadElement() ReaderElement {
	defer func() { v.currentIndex++ }()

	if v.currentIndex >= len(v.sourceStrings) {
		return NewFinalReaderElement()
	}
	return NewReaderElement(v.sourceStrings[v.currentIndex])
}
