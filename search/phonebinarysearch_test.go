/*
@Time : 2019/8/19 15:35
@Author : louis
@File : phoneBinarySearch_test
@Software: GoLand
*/

package search

import (
	"fmt"
	"testing"
	"time"
)

func since(t time.Time) int64 {
	return time.Since(t).Nanoseconds()
}

func init() {
	Debug()
}

func BenchmarkFind(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var i = 0
		for pb.Next() {
			i++
			_, err := Find(fmt.Sprintf("%s%d%s", "1897", i&10000, "45"))
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func TestFindPhone1(t *testing.T) {
	_, err := Find("123412325322213")
	if err != nil {
		t.Fatal("err result",err)
	}
	t.Log(err)
}

func TestFindPhone2(t *testing.T) {
	_, err := Find("15821586468")
	if err != nil {
		t.Fatal("err result",err)
	}
	t.Log(err)
}

func TestFindPhone3(t *testing.T) {
	_, err := Find("1703576")
	if err != nil {
		t.Fatal("err result",err)
	}
	t.Log(err)
}

func TestFindPhone4(t *testing.T) {
	_, err := Find("13274872323")
	if err != nil {
		t.Fatal("err result",err)
	}
	t.Log(err)
}

func TestFindPhone5(t *testing.T) {
	_, err := Find("bdsa32323")
	if err != nil {
		t.Fatal("err result",err)
	}
	t.Log(err)
}
