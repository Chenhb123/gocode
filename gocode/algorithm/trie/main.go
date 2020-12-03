package main

import "fmt"

// 前缀树

/*
Trie树，即字典树，又称单词查找树或键树，是一种树形结构，是一种哈希树的变种。
典型应用是用于统计和排序大量的字符串（但不仅限于字符串），所以经常被搜索引擎系统用于文本词频统计。
它的优点是：最大限度地减少无谓的字符串比较，查询效率比哈希表高。
小写英文字母或大写英文字母的字典数是一个26叉树。Trie树的根结点是不保存数据的，所有的数据都保存在它的孩子节点中.
*/

// MaxCap 最大叶子节点个数
var MaxCap = 26 // a-z

// Trie 前缀树结构
type Trie struct {
	Next   map[rune]*Trie
	IsWord bool
}

// Constructor 构造器 initialzie your structure here
func Constructor() Trie {
	root := new(Trie)
	root.Next = make(map[rune]*Trie, MaxCap)
	root.IsWord = false
	return *root
}

// Insert insert a word into trie
func (t *Trie) Insert(word string) {
	for _, v := range word {
		if t.Next[v] == nil {
			node := new(Trie)
			node.Next = make(map[rune]*Trie, MaxCap)
			node.IsWord = false
			t.Next[v] = node
		}
		t = t.Next[v]
	}
	t.IsWord = true
}

// Search returns true if word in the trie
func (t *Trie) Search(word string) bool {
	for _, v := range word {
		if t.Next[v] == nil {
			return false
		}
		t = t.Next[v]
	}
	return t.IsWord
}

// StartWith returns true if any word in the trie that starts with the given prefix
func (t *Trie) StartWith(prefix string) bool {
	for _, v := range prefix {
		if t.Next[v] == nil {
			return false
		}
		t = t.Next[v]
	}
	return true
}

func main() {
	t := Constructor()
	t.Insert("hello")
	fmt.Println(t.Search("hello"))
	fmt.Println(t.StartWith("hee"))
}
