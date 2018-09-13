//Copyright Copyright 2018 Andrey Z.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0

package trie

func (t *Node) IsFinal() bool {
	return t.isFinal
}

func (t *Node) Children() []INode {
	ch := make([]INode, 0)
	for _, c := range t.children {
		ch = append(ch, c)
	}
	return ch
}

func (t *Node) Symbol() rune {
	return t.letter
}

func (t *Node) Words() []string {
	res := make([]string, 0)
	for _, n := range t.children {
		if n.isFinal {
			res = append(res, string([]rune{t.letter, n.letter}))
		}

		for _, w := range n.Words() {
			newWord := append([]rune{t.letter}, []rune(w)...)
			res = append(res, string(newWord))
		}
	}
	return res
}

func (t *Node) Put(word string) {
	node := t
	for _, char := range []rune(word) {
		if node.children[char] == nil {
			newNode := newNode()
			newNode.letter = char

			node.children[char] = newNode
			node = newNode
		} else {
			node = node.children[char]
		}
	}
	node.isFinal = true
}

func (t *Node) Wrap() INode {
	return &Node{children: map[rune]*Node{t.letter: t}}
}
