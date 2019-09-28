/*
Copyright 2019 louis.
@Time : 2019/9/27 15:56
@Author : louis
@File : 133-longestword
@Software: GoLand

*/

/*
### **Description**


给一个词典，找出其中所有最长的单词。

### **Example**

```
样例 1:
	输入:   {
		"dog",
		"google",
		"facebook",
		"internationalization",
		"blabla"
		}
	输出: ["internationalization"]



样例 2:
	输入: {
		"like",
		"love",
		"hate",
		"yes"
		}
	输出: ["like", "love", "hate"]
```

### **Challenge**

遍历两次的办法很容易想到，如果只遍历一次你有没有什么好办法？
*/

package lintcode

func LongestWord(m []string) []string {
	var res []string
	max := 0
	for _, v := range m {
		if len(v) > max {
			max = len(v)
			res = []string{} //大于则清空数组,将最大的加入到数组
			res = append(res, v)
			//len 等于最大的,则直接append至数组.
		} else if len(v) == max {
			res = append(res, v)
		}
	}
	return res
}
