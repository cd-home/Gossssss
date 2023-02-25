package main

import (
	"sync"
)

type LazySingleton struct{}

var lazySingleton *LazySingleton
var once sync.Once

// GetLazyInstance while use Instance will to initialize lazySingleton
func GetLazyInstance() *LazySingleton {
	if lazySingleton != nil {
		once.Do(func() {
			lazySingleton = &LazySingleton{}
		})
	}
	return lazySingleton
}

func main() {

}
