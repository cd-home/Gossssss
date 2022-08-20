package main

import "fmt"

type Trie struct {
	isWord   bool
	next map[string]*Trie
}

func Constructor() Trie {
	return Trie{
		isWord: false,
		next: make(map[string]*Trie),
	}
}

func (t *Trie) Insert(words string) {
	cur := t
	for _, word := range words {
		node, ok := cur.next[string(word)]
		if !ok {
			child := &Trie{}
			child.next = make(map[string]*Trie)
			child.isWord = false
			cur.next[string(word)] = child
		} else {
			cur = node
		}
	}
	cur.isWord = true
}

func (t *Trie) Search(words string) bool {
	cur := t
	for _, word := range words {
		_, ok := cur.next[string(word)]
		if !ok {
			return false
		}
	}
	return cur.isWord
}

func (t *Trie) StartWith(prefix string) bool {
	cur := t
	for _, word := range prefix {
		fmt.Println(string(word))
		node, ok := cur.next[string(word)]
		if !ok {
			return false
		}
		cur = node
	}
	return true
}

func main() {
	//for _, v := range "你好" {
	//	fmt.Println(string(v))
	//	fmt.Println(v)
	//}
	//testString := "你好，世界"
	//fmt.Println(testString[:2]) // 输出乱码，因为截取了前两个字节
	//fmt.Println(testString[:3]) // 输出「你」，一个中文字符由三个字节表示
	s1 := "你好12314"
	s2 := "你好吗"
	s3 := "淘宝镜像"
	s4 := "你好，世界"

	t := Constructor()
	t.Insert(s1)
	t.Insert(s2)
	t.Insert(s3)
	t.Insert(s4)

	//fmt.Println(t.Search("你好"))
	fmt.Println(t.StartWith("你好吗"))
	fmt.Println(t)
}
