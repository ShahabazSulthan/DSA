package main

import (
	"fmt"
	"strings"
)

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
	}
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func (t *Trie) insert(word string) {
	currentNode := t.root
	for _, char := range word {
		if _, exist := currentNode.children[char]; !exist {
			currentNode.children[char] = NewTrieNode()
		}
		currentNode = currentNode.children[char]
	}
	currentNode.isEnd = true
}

func (t *Trie) Search(words string) bool {
	currentNode := t.root
	for _, char := range words {
		if _, exist := currentNode.children[char]; !exist {
			return false
		}
		currentNode = currentNode.children[char]
	}
	return currentNode.isEnd
}

func (t *Trie) StartsWith(word string) bool {
	currentNode := t.root
	for _, char := range word {
		if _, exist := currentNode.children[char]; !exist {
			return false
		}
		currentNode = currentNode.children[char]
	}
	return true
}

func main() {
	trie := NewTrie()
	words := []string{"hello", "world", "trie", "tree", "algo"}

	for _, word := range words {
		trie.insert(strings.ToLower(word))
	}

	// Search words in Trie

	fmt.Println("Search words in Trie")
	fmt.Println("")
	fmt.Println(trie.Search("hello"))
	fmt.Println(trie.Search("world"))
	fmt.Println(trie.Search("trie"))
	fmt.Println(trie.Search("tree"))
	fmt.Println(trie.Search("algo"))
	fmt.Println(trie.Search("algos"))

	fmt.Println("")
	// Check if there are any words starting with the given prefix
	fmt.Println("Check if there are any words starting with the given prefix")
	fmt.Println(trie.StartsWith("tr"))
	fmt.Println(trie.StartsWith("he"))
	fmt.Println(trie.StartsWith("wo"))
	fmt.Println(trie.StartsWith("al"))
	fmt.Println(trie.StartsWith("xx"))
}
