/*
@Time : 2019/9/19 23:26
@Author : louis
@File : btree_test
@Software: GoLand
*/

package tree

import (
	"fmt"
	"testing"
)

func TestBtree(t *testing.T) {
	b := &btree{}
	b.NewBtree()
	b.PreOrder()
	fmt.Println()
	b.PostOrder()
	fmt.Println()
	b.InOrder()
	fmt.Println()
	b.LevelOrder()
	fmt.Println()
}

func BenchmarkBtree_InOrder(b *testing.B) {
	bt := &btree{}
	bt.NewBtree()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bt.InOrder()
	}
}
func BenchmarkBtree_PreOrder(b *testing.B) {
	bt := &btree{}
	bt.NewBtree()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bt.PreOrder()
	}
}
func BenchmarkBtree_PostOrder(b *testing.B) {
	bt := &btree{}
	bt.NewBtree()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bt.PostOrder()
	}
}

func BenchmarkBtree_LevelOrder(b *testing.B) {
	bt := &btree{}
	bt.NewBtree()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bt.LevelOrder()
	}
}