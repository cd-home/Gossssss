package leetcodesss

import (
	"container/list"
	"errors"
)

type CacheNode struct {
	Key, Value interface{}
}

func (node *CacheNode) NewCacheNode(k, v interface{}) *CacheNode {
	return &CacheNode{Key:k, Value:v}
}

type LRUCache struct {
	Cap int
	doubleList *list.List
	CacheMap map[interface{}]*list.Element
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		Cap:cap,
		doubleList:list.New(),
		CacheMap: make(map[interface{}]*list.Element),
	}
}

func (lru *LRUCache) Size() int {
	return lru.doubleList.Len()
}

func (lru *LRUCache) Set(k, v interface{}) error {
	// 判断链表
	if lru.doubleList == nil {
		return errors.New("can not init lru")
	}
	// 如果在缓存中
	if pElement, ok := lru.CacheMap[k]; ok {
		pElement.Value.(*CacheNode).Value = v
		lru.doubleList.MoveToFront(pElement)
		return nil
	}
	// 不在缓存中
	newElement := lru.doubleList.PushFront(&CacheNode{k, v})
	lru.CacheMap[k] = newElement
	// 容量超出
	if lru.doubleList.Len() > lru.Cap {
		lastElement := lru.doubleList.Back()
		if lastElement == nil {
			return nil
		}
		cacheNode := lastElement.Value.(*CacheNode)
		delete(lru.CacheMap, cacheNode.Key)
		lru.doubleList.Remove(lastElement)
	}
	return nil
}

func (lru *LRUCache) Get(k interface{}) (v interface{}, ret bool, err error) {
	if lru.CacheMap == nil {
		return v, false, errors.New("failed to init lru")
	}
	if pElement, ok := lru.CacheMap[k]; ok {
		lru.doubleList.MoveToFront(pElement)
		return pElement.Value.(*CacheNode).Value, true, nil
	}
	return v, false, nil
}

func (lru *LRUCache) Remove(k interface{}) bool {
	if lru.CacheMap == nil {
		return false
	}
	if pElement, ok := lru.CacheMap[k]; ok {
		cacheNode := pElement.Value.(*CacheNode)
		delete(lru.CacheMap, cacheNode.Key)
		lru.doubleList.Remove(pElement)
		return true
	}
	return false
}