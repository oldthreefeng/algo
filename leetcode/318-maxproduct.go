/*
Copyright 2019 louis.
@Time : 2019/10/8 0:33
@Author : louis
@File : 318-maxproduct
@Software: GoLand

*/

package leetcode

/*
String2int 思路:
用二进制的一位表示某一个字母是否出现过，0表示没出现，1表示出现。
"abcd"二进制表示00000000 00000000 00000000 00001111,
"bc"二进制表示00000000 00000000 00000000 00000110。
当两个字符串没有相同的字母时，二进制数与的结果为0。
*/
func String2int(str string) (res int) {
	for i := 0; i < len(str); i++ {  //不能用for-range迭代.
		res |= 1 << uint(str[i]-'a')   // "abc" ==> 二进制"111"
	}
	return res
}

func MaxProduct(words []string) int {
	var arr []int = make([]int, len(words))
	l := len(words)
	for i := 0; i < l; i++ {
		arr[i] = String2int(words[i])  //将[]string数组里面的string,转化为int.
	}
	var res int
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			// 遍历数组,如果数组里面&操作为0,即这两个单词不含有公共字母
			// 并且res < length(word[i]) * length(word[j])时,更新res.
			if arr[i]&arr[j] == 0 && len(words[i])*len(words[j]) > res {
				res = len(words[i]) * len(words[j])
			}
		}
	}
	return res
}
