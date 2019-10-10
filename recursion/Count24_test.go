/**
 * @time: 2019-08-20 22:24
 * @author: louis
 */
package recursion

import (
	"testing"
)

func TestJudgePoint24(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test01", args{[]int{6, 6, 6, 6}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JudgePoint24(tt.args.nums); got != tt.want {
				t.Errorf("JudgePoint24() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCount24(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test01", args{[]int{1,2,1,2}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Count24(tt.args.nums); got != tt.want {
				t.Errorf("Count24() = %v, want %v", got, tt.want)
			}
		})
	}
}
