/*
@Time : 2019/10/9 11:24
@Author : louis
@File : reverse_test.go
@Software: GoLand
*/

package utils

import (
	"testing"
)

func BenchmarkReverse(b *testing.B) {
	var s = "abcsdasda321dsad3"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reverse(s)
	}
}

func BenchmarkReverseX(b *testing.B) {
	var s = "abcsdasda321dsad3"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReverseX(s)
	}
}

func BenchmarkReverseXa(b *testing.B) {
	var s = "abcsdasda321dsad3"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReverseXa(s)
	}
}

func TestReverseXa(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test01", args{"abcsdasda321dsad3"}, "3dasd123adsadscba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseXa(tt.args.s); got != tt.want {
				t.Errorf("ReverseXa() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseX(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test01", args{"abcsdasda321dsad3"}, "3dasd123adsadscba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseX(tt.args.s); got != tt.want {
				t.Errorf("ReverseX() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
$ go test -bench=. -benchmem -run=none
goos: windows
goarch: amd64
pkg: gogs.wangke.co/go/algo/utils
BenchmarkReverse-4                       3833876               309 ns/op              32 B/op          1 allocs/op
BenchmarkReverseX-4                      4060528               275 ns/op              32 B/op          1 allocs/op
BenchmarkReverseXa-4                     4213195               345 ns/op             112 B/op          2 allocs/op

*/