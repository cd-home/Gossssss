package main

import (
	"fmt"
)

const SIZE = 15

type Node struct {
	Value int
	Next  *Node
}

type HashTable struct {
	Table map[int]*Node
	Size  int
}

func hashFunction(i, size int) int {
	return i % size
}

func insertV(hash *HashTable, value int) int {
	// 先找桶
	index := hashFunction(value, hash.Size)
	e := Node{
		Value: value,
		Next:  hash.Table[index],
	}
	hash.Table[index] = &e
	return index
}

func traverseHash(hash *HashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%d -> ", t.Value)
				t = t.Next
			}
		}
		fmt.Println()
	}
}

func lookUp(hash *HashTable, value int) bool {
	index := hashFunction(value, hash.Size)
	if hash.Table[index] != nil {
		t := hash.Table[index]
		for t != nil {
			if t.Value == value {
				return true
			}
			t = t.Next
		}
	}
	return false
}

func main() {
	table := make(map[int]*Node, SIZE)
	hash := &HashTable{
		Table: table,
		Size:  SIZE,
	}
	for i := 0; i < 120; i++ {
		insertV(hash, i)
	}
	traverseHash(hash)
}
