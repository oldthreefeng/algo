/**
 * @time: 2019-08-21 01:09
 * @author: louis
 */
package main

import (
	"fmt"
	"gogs.wangke.co/go/algo/recursion"
)

func main() {
	var a = []int{6, 6, 6, 6}
	if recursion.Count24(a) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
