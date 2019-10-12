/*
@Time : 2019/10/12 13:40
@Author : louis
@File : 146-lrucache_test.go
@Software: GoLand
*/

package leetcode

import (
	"fmt"
	"testing"
)

func TestConstructorLRU(t *testing.T) {
	var cache = NewLRU(2)
	cache.Put(1,1)
	cache.Put(2,2)
	fmt.Println(cache.Get(1))
	cache.Put(3,3)
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(3))
	cache.Put(4,4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
	//1 -1 3 -1 3 4
}