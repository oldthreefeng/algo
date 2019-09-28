/*
Copyright 2019 louis.
@Time : 2019/9/28 16:42
@Author : louis
@File : 35-reverse
@Software: GoLand

*/

package lintcode

import (
	"reflect"
	"testing"
)


/*
### **Example**

**样例 1:**

```
输入: 1->2->3->null
输出: 3->2->1->null
```

**样例 2:**

```
输入: 1->2->3->4->null
输出: 4->3->2->1->null
```
*/
func TestReverseLink(t *testing.T) {
	type args struct {
		l *LinkNode
	}
	tests := []struct {
		name string
		args args
		want *LinkNode
	}{
		{"test01",args{&LinkNode{1,&LinkNode{2,&LinkNode{3,nil}}}},&LinkNode{3,&LinkNode{2,&LinkNode{1,nil}}}},
		{"test02",args{&LinkNode{1,&LinkNode{2,&LinkNode{3,&LinkNode{4,nil}}}}},&LinkNode{4,&LinkNode{3,&LinkNode{2,&LinkNode{1,nil}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseLink(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkReverseLink(b *testing.B) {
	l1 := &LinkNode{1,&LinkNode{2,&LinkNode{3,&LinkNode{4,nil}}}}
	b.ResetTimer()
	for i:=0;i<b.N;i++ {
		ReverseLink(l1)
	}
}

func BenchmarkReverseLinkX(b *testing.B) {
	l1 := &LinkNode{1,&LinkNode{2,&LinkNode{3,&LinkNode{4,nil}}}}
	b.ResetTimer()
	for i:=0;i<b.N;i++ {
		ReverseLinkX(l1)
	}
}

/*
goos: windows
goarch: amd64
pkg: gogs.wangke.co/go/algo/lintcode
BenchmarkReverseLink-4    	289194604	         3.96 ns/op
BenchmarkReverseLinkX-4   	255333871	         4.75 ns/op
PASS
利用空间换时间,优化了1 ns/op
*/