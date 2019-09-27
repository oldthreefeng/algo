/*
Copyright 2019 louis.
@Time : 2019/9/27 15:56
@Author : louis
@File : 133-longestword
@Software: GoLand

*/

package lintcode

import (
	"gogs.wangke.co/go/algo/utils"
	"testing"
)

func TestLongestWord(t *testing.T) {
	type args struct {
		m []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"test01",args{[]string{"dog",
			"google",
			"facebook",
			"internationalization",
			"blabla",}},[]string{"internationalization"}},
		{"test02",args{[]string{"like",
			"love",
			"hate",
			"yes",}},[]string{"like", "love", "hate"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestWord(tt.args.m); !utils.StringSliceEqualBce(got, tt.want) {
				t.Errorf("LongestWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

