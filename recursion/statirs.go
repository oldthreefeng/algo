/**
 * @time: 2019-08-20 22:01
 * @author: louis
 */
package recursion

func Stairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return Stairs(n-2) + Stairs(n-1)
}
