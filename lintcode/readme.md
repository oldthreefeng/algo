# lintcode 日常编程golang

Table of Contents
=================

   * [lintcode 日常编程golang](#lintcode-日常编程golang)
      * [02-trailingzeros](#02-trailingzeros)
      * [03-digitcounts](#03-digitcounts)
      * [04-uglynumber](#04-uglynumber)
      * [05-Kthlargestelement](#05-kthlargestelement)
         * [<strong>Example</strong>](#example)
         * [<strong>Challenge</strong>](#challenge)
      * [06-mergesortedArray](#06-mergesortedarray)
      * [08-rotatingString](#08-rotatingstring)
      * [09-fizzbuzz](#09-fizzbuzz)
      * [11-searchRange](#11-searchrange)
      * [13-strstr](#13-strstr)
         * [<strong>Description</strong>](#description)
         * [<strong>Clarification</strong>](#clarification)
         * [<strong>Example</strong>](#example-1)
         * [<strong>Challenge</strong>](#challenge-1)
      * [14-binarysearch](#14-binarysearch)
         * [<strong>Description</strong>](#description-1)
         * [<strong>Example</strong>](#example-2)
         * [<strong>Challenge</strong>](#challenge-2)
      * [15-permute](#15-permute)
      * [28-searcMatrix](#28-searcmatrix)
         * [Description](#description-2)
         * [<strong>Example</strong>](#example-3)
         * [<strong>Challenge</strong>](#challenge-3)
      * [29-patitionArray](#29-patitionarray)
         * [Example](#example-4)
         * [Challenge](#challenge-4)
         * [Notice](#notice)
      * [35-reverse](#35-reverse)
         * [<strong>Example</strong>](#example-5)
      * [49-sortletters](#49-sortletters)
         * [<strong>Description</strong>](#description-3)
         * [<strong>Example</strong>](#example-6)
         * [<strong>Challenge</strong>](#challenge-5)
      * [54-atoi](#54-atoi)
         * [<strong>Description</strong>](#description-4)
         * [<strong>Example</strong>](#example-7)
      * [56-Twosum](#56-twosum)
         * [<strong>Description</strong>](#description-5)
         * [<strong>Example</strong>](#example-8)
      * [128-hashcode](#128-hashcode)
      * [129-rehashing](#129-rehashing)
      * [133-lonestword](#133-lonestword)
         * [<strong>Description</strong>](#description-6)
         * [<strong>Example</strong>](#example-9)
         * [<strong>Challenge</strong>](#challenge-6)
      * [135-combinationsum](#135-combinationsum)
         * [<strong>Description</strong>](#description-7)
         * [<strong>Example</strong>](#example-10)
      * [1324-countprimes](#1324-countprimes)

Created by [gh-md-toc](https://github.com/ekalinin/github-markdown-toc)

## 02-trailingzeros

```cgo
设计一个算法，计算出n阶乘中尾部零的个数

Example
样例  1:
	输入: 11
	输出: 2

	样例解释:
	11! = 39916800, 结尾的0有2个。

样例 2:
	输入:  5
	输出: 1

	样例解释:
	5! = 120， 结尾的0有1个。

Challenge
O(logN)的时间复杂度
```
[传送门](./02-trailingzeros.go)
## 03-digitcounts

```cgo
计算数字 k 在 0 到 n 中的出现的次数，k 可能是 0~9 的一个值。

Example
样例 1：

输入：
k = 1, n = 1
输出：
1
解释：
在 [0, 1] 中，我们发现 1 出现了 1 次 (1)。
样例 2：

输入：
k = 1, n = 12
输出：
5
解释：
在 [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12] 中，我们发现 1 出现了 5 次 (1, 10, 11, 12)(注意11中有两个1)。

```
[传送门](./03-digitcounts.go)
## 04-uglynumber

```cgo
设计一个算法，找出只含素因子2，3，5 的第 n 小的数。

符合条件的数如：1, 2, 3, 4, 5, 6, 8, 9, 10, 12...

Example
样例 1：

输入：9
输出：10
样例 2：

输入：1
输出：1
Challenge
要求时间复杂度为 O(nlogn) 或者 O(n)。

Notice
我们可以认为 1 也是一个丑数。
```
[传送门](./04-uglynumber.go)
## 05-Kthlargestelement

### **Example**

**样例 1：**

```
输入：
n = 1, nums = [1,3,4,2]
输出：
4
```

**样例 2：**

```
输入：
n = 3, nums = [9,3,2,4,8]
输出：
4
```

### **Challenge**

要求时间复杂度为O(n)，空间复杂度为O(1)。
[传送门](./05-kthlargestelement.go)
## 06-mergesortedArray

```cgo
合并排序数组 II
中文English
合并两个有序升序的整数数组A和B变成一个新的数组。新数组也要有序。

Example
样例 1:

输入: A=[1], B=[1]
输出:[1,1]
样例解释: 返回合并后的数组。
样例 2:

输入: A=[1,2,3,4], B=[2,4,5,6]
输出: [1,2,2,3,4,4,5,6]
样例解释: 返回合并后的数组。
Challenge
你能否优化你的算法，如果其中一个数组很大而另一个数组很小？
```
[传送门](./06-mergeSortedArray.go)
## 08-rotatingString

```cgo
给定一个字符串（以字符数组的形式给出）和一个偏移量，根据偏移量原地旋转字符串(从左向右旋转)。

Example
样例 1:

输入:  str="abcdefg", offset = 3
输出:  str = "efgabcd"
样例解释:  注意是原地旋转，即str旋转后为"efgabcd"
样例 2:

输入: str="abcdefg", offset = 0
输出: str = "abcdefg"
样例解释: 注意是原地旋转，即str旋转后为"abcdefg"
样例 3:

输入: str="abcdefg", offset = 1
输出: str = "gabcdef"
样例解释: 注意是原地旋转，即str旋转后为"gabcdef"
样例 4:

输入: str="abcdefg", offset =2
输出: str = "fgabcde"
样例解释: 注意是原地旋转，即str旋转后为"fgabcde"
样例 5:

输入: str="abcdefg", offset = 10
输出: str = "efgabcd"
样例解释: 注意是原地旋转，即str旋转后为"efgabcd"
Challenge
在数组上原地旋转，使用O(1)的额外空间
```
[传送门](./08-rotateString.go)
## 09-fizzbuzz

```cgo
给你一个整数n. 从 1 到 n 按照下面的规则打印每个数：

如果这个数被3整除，打印fizz.
如果这个数被5整除，打印buzz.
如果这个数能同时被3和5整除，打印fizz buzz.
如果这个数既不能被 3 整除也不能被 5 整除，打印数字本身。
Example
比如 n = 15, 返回一个字符串数组：

[
  "1", "2", "fizz",
  "4", "buzz", "fizz",
  "7", "8", "fizz",
  "buzz", "11", "fizz",
  "13", "14", "fizz buzz"
]
Challenge
你是否可以只用一个 if 来实现

```
[传送门](./09-fizzbuzz.go)
## 11-searchRange

```cgo
给定一个二叉查找树和范围[k1, k2]。按照升序返回给定范围内的节点值。

Example
样例 1:

输入：{5},6,10
输出：[]
        5
它将被序列化为 {5}
没有数字介于6和10之间
样例 2:

输入：{20,8,22,4,12},10,22
输出：[12,20,22]
解释：
        20
       /  \
      8   22
     / \
    4   12
它将被序列化为 {20,8,22,4,12}
[12,20,22]介于10和22之间
```
[传送门](./11-searchRange.go)
## 13-strstr

### **Description**

对于一个给定的 source 字符串和一个 target 字符串，你应该在 source 字符串中找出 target 字符串出现的第一个位置(从0开始)。如果不存在，则返回 `-1`。


### **Clarification**

在面试中我是否需要实现KMP算法？

- 不需要，当这种问题出现在面试中时，面试官很可能只是想要测试一下你的基础应用能力。当然你需要先跟面试官确认清楚要怎么实现这个题。

### **Example**

**样例 1:**

```
输入: source = "source" ， target = "target"
输出:-1	
样例解释: 如果source里没有包含target的内容，返回-1
```

**样例 2:**

```
输入: source = "abcdabcdefg" ，target = "bcd"
输出: 1	
样例解释: 如果source里包含target的内容，返回target在source里第一次出现的位置
```

### **Challenge**

O(n2)的算法是可以接受的。如果你能用O(n)的算法做出来那更加好。（提示：KMP）
[传送门](./13-strStr.go)
## 14-binarysearch

### **Description**

中文

给定一个排序的整数数组（升序）和一个要查找的整数`target`，用`O(logn)`的时间查找到target第一次出现的下标（从0开始），如果target不存在于数组中，返回`-1`。

### **Example**

```
样例  1:
	输入:[1,4,4,5,7,7,8,9,9,10]，1
	输出: 0
	
	样例解释: 
	第一次出现在第0个位置。

样例 2:
	输入: [1, 2, 3, 3, 4, 5, 10]，3
	输出: 2
	
	样例解释: 
	第一次出现在第2个位置
	
样例 3:
	输入: [1, 2, 3, 3, 4, 5, 10]，6
	输出: -1
	
	样例解释: 
	没有出现过6， 返回-1
```

### **Challenge**

如果数组中的整数个数超过了2^32，你的算法是否会出错？
[传送门](./14-binarysearch.go)
## 15-permute

//TODO
[传送门](./15-permute.go)
## 28-searcMatrix

### Description

写出一个高效的算法来搜索 *m* × *n*矩阵中的值。

这个矩阵具有以下特性：

- 每行中的整数从左到右是排序的。
- 每行的第一个数大于上一行的最后一个整数。

### **Example**

```
样例  1:
	输入: [[5]],2
	输出: false

	样例解释:
  没有包含，返回false。

样例 2:
	输入:
[
  [1, 3, 5, 7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
],3
	输出: true

	样例解释:
	包含则返回true。
```

### **Challenge**

O(log(n) + log(m)) 时间复杂度
[传送门](./28-searchMatrix.go)
## 29-patitionArray

给出一个整数数组 *nums* 和一个整数 *k*。划分数组（即移动数组 *nums* 中的元素），使得：

- 所有小于k的元素移到左边
- 所有大于等于k的元素移到右边

返回数组划分的位置，即数组中第一个位置 *i*，满足 *nums*[*i*] 大于等于 *k*。

### Example

例1:

```
输入:
[],9
输出:
0
```

例2:

```
输入:
[3,2,2,1],2
输出:1
解释:
真实的数组为[1,2,2,3].所以返回 1
```

### Challenge

使用 O(n) 的时间复杂度在数组上进行划分。

### Notice

你应该真正的划分数组 *nums*，而不仅仅只是计算比 *k* 小的整数数，如果数组 *nums* 中的所有元素都比 *k* 小，则返回 nums.length。

[传送门](./29-partitionArray.go)
## 35-reverse

### **Example**

**样例 1:**

```
输入: 1->2->3->null
输出: 3->2->1->null
```

**样例 2:**

```
输入: 1->2->3->4->null
输出: 4->3->2->1->null
```
[传送门](./35-reverse.go)
## 49-sortletters

### **Description**

给定一个只包含字母的字符串，按照先小写字母后大写字母的顺序进行排序。

小写字母或者大写字母他们之间不一定要保持在原始字符串中的相对位置。

### **Example**

```
样例 1:
	输入:  "abAcD"
	输出:  "acbAD"

样例 2:
	输入: "ABC"
	输出:  "ABC"
	
```

### **Challenge**

在原地扫描一遍完成
[传送门](./49-sortLetters.go)
## 54-atoi

### **Description**


实现atoi这个函数，将一个字符串转换为整数。如果没有合法的整数，返回0。如果整数超出了32位整数的范围，返回INT_MAX(2147483647)如果是正整数，或者INT_MIN(-2147483648)如果是负整数。

### **Example**

**样例1**

```plain
输入: "10"
输出: 10
```

**样例2**

```plain
输入: "1"
输出: 1
```

**样例3**

```plain
输入: "123123123123123"
输出: 2147483647
说明: 因为 123123123123123 > INT_MAX, 所以我们返回INT_MAX
```

**样例4**

```plain
输入: "1.0"
输出: 1
说明: 我们只需要转换第一个有效数字
```
[传送门](./54-atoi.go)
## 56-Twosum

### **Description**

中文

给一个整数数组，找到两个数使得他们的和等于一个给定的数 *target*。

你需要实现的函数`twoSum`需要返回这两个数的下标, 并且第一个下标小于第二个下标。注意这里下标的范围是 0 到 *n-1*。

你可以假设只有一组答案。
### **Example**

```
Example1:
给出 numbers = [2, 7, 11, 15], target = 9, 返回 [0, 1].
Example2:
给出 numbers = [15, 2, 7, 11], target = 9, 返回 [1, 2].
```
[传送门](./56-Twosum.go)
## 128-hashcode
[传送门](./128-hashcode.go)
## 129-rehashing
[传送门](./129-rehashing.go)

## 133-lonestword
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
[传送门](./133-longestword.go)
## 135-combinationsum

### **Description**

中文English

给定一个候选数字的集合 `candidates` 和一个目标值 `target`. 找到 `candidates` 中所有的和为 `target` 的组合.

在同一个组合中, `candidates` 中的某个数字不限次数地出现.

所有数值 (包括 `target` ) 都是正整数.返回的每一个组合内的数字必须是非降序的.返回的所有组合之间可以是任意顺序.解集不能包含重复的组合.


### **Example**

**样例 1:**

```
输入: candidates = [2, 3, 6, 7], target = 7
输出: [[7], [2, 2, 3]]
```

**样例 2:**

```
输入: candidates = [1], target = 3
输出: [[1, 1, 1]]
```
[传送门](./135-combinationsum.go)
## 1324-countprimes

```cgo
计算小于非负数n的质数的个数。

Example
样例 1

输入: n = 2
输出: 0
样例 2

输入: n = 4
输出: 2
解析：2, 3 是素数
```
[传送门](./1324-CountPrimes.go)