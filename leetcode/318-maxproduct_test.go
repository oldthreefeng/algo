/*
Copyright 2019 louis.
@Time : 2019/10/8 0:33
@Author : louis
@File : 318-maxproduct
@Software: GoLand

*/

package leetcode

import (
	"testing"
)

func TestMaxProduct(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test01", args{[]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}}, 4},
		{"test02", args{[]string{"abcw","baz","foo","bar","xtfn","abcdef"}}, 16},
		{"test03", args{[]string{"a","aa","aaa","aaaa"}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProduct(tt.args.words); got != tt.want {
				t.Errorf("MaxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString2int(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"test01",args{"abcd"},0x0f},
		{"test02",args{"bc"},0x06},
		{"test03",args{"abcdefghijklmnopqrstuvwxyz"},0x3ffffff},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := String2int(tt.args.str); gotRes != tt.wantRes {
				t.Errorf("String2int() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
