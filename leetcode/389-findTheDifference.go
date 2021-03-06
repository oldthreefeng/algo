/*
@Time : 2019/10/9 18:54
@Author : louis
@File : 389-findTheDifference
@Software: GoLand
*/

package leetcode

/*
FindTheDifference
Given two strings s and t which consist of only lowercase letters.
String t is generated by random shuffling string s and then add one more letter at a random position.
Find the letter that was added in t.
*/
func FindTheDifference(s, t string) byte {
	var res byte
	for i := 0; i < len(s); i++ {
		res ^= s[i]
	}
	for i := 0; i < len(t); i++ {
		res ^= t[i]
	}
	return res
}
