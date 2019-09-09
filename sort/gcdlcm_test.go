/*
@Time : 2019/9/9 10:18
@Author : louis
@File : gcdlcm_test
@Software: GoLand
*/

package sort

import (
	"fmt"
	"testing"
)

func TestGcd(t *testing.T) {
	a := Gcd(35, 215)
	b := Lcm(49, 35)
	fmt.Println(a, b) // 5 245
}
