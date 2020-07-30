/*
@Time : 2019/10/8 13:57
@Author : louis
@File : 5214-longestsubsequence
@Software: GoLand
*/

package leetcode

import "github.com/oldthreefeng/algo/utils"

/*
LongestSubsequence 最长等差子序列
这道题思路比较简单，跟经典问题`最长递增（减）子序列`有点相似，而这道题称为`最长等差子序列`, 并且公差还是固定的，相对还更简单一点。
可以用`dp[i]`来记录以数字`i`为结尾的`最长等差子序列`的长度，那么它应该只有两种情况：

- `dp[i] = 1 // 表示在 i 之前没有出现等差子序列`;

- `dp[i] = dp[i - difference] + 1 // 表示在 i 之前出现了等差子序列，长度为 dp[i-difference], 而 i 也是满足这个等差序列的，所以等差序列的长度在此基础上加 1 就可以了`

考虑元素值会出现负数，所以用数组存放是不行的，那么可以用一个 `map`来维护以 `i` 结尾的最长等差序列的长度，所以也就不难得出如下代码

设m[i]表示当前序列最后一个数是i的最大子序列长度
递推式: m[i]=max(m[i],m[i-difference]+1);
*/
func LongestSubsequence(arr []int, difference int) int {
	m := make(map[int]int, len(arr))
	var ans int
	for i := 0; i < len(arr); i++ {
		sum := arr[i] - difference
		m[arr[i]] = utils.Max(m[arr[i]], m[sum]+1)
		ans = utils.Max(ans, m[arr[i]])
	}
	return ans
}
