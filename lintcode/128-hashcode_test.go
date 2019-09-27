/*
Copyright 2019 louis.
@Time : 2019/9/27 14:42
@Author : louis
@File : 128-hashcode
@Software: GoLand

*/

package lintcode

import "testing"

func TestHashCode(t *testing.T) {
	type args struct {
		key       string
		HASH_SIZE int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1",args{"abcd",1000},978,},
		{"test2",args{"abcd",100},78,},
		{"test3",args{"abcdefghijklmnopqrstuvwxyz",2607},1673,},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HashCode(tt.args.key, tt.args.HASH_SIZE); got != tt.want {
				t.Errorf("HashCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
