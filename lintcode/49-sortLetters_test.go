/*
Copyright 2019 louis.
@Time : 2019/9/28 23:14
@Author : louis
@File : 49-sortLetters
@Software: GoLand

*/

package lintcode

import "testing"

func TestSortLetters(t *testing.T) {
	type args struct {
		st string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test01",args{"abAcD"},"abcAD",},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortLetters(tt.args.st); got != tt.want {
				t.Errorf("SortLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
