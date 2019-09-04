/*
@Time : 2019/9/4 14:16
@Author : louis
@File : poly_test
@Software: GoLand
*/

package main

import (
	"fmt"
	"testing"
)

func TestCreatePoly(t *testing.T) {
	s := &PolyNode{}
	fmt.Scanln(&s.coef,s.exp)
	fmt.Print(s.coef,s.exp)
}
