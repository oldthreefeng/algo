/**
 * @time: 2019-08-20 00:14
 * @author: louis
 */
package recursion

import "fmt"

func Hanoi(n int, src, mid, dest string) {
	if n == 1 {
		fmt.Printf("%s -> %s\n", src, dest)
		return
	}
	Hanoi(n-1, src, dest, mid) // 把n-1的src移动到mid,相当于中转站
	fmt.Printf("%s -> %s\n", src, dest)
	Hanoi(n-1, mid, src, dest) //把n-1的mid移动到dest
}
