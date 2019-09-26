/*
@Time : 2019/9/26 15:23
@Author : louis
@File : 11-searchRange
@Software: GoLand
*/

package lintcode

/*
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
*/


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
 

func searchRange (root *TreeNode, k1 int, k2 int) []int {
	// write your code here
	res := make([]int,0)
	InOrder(root,&res,k1,k2)
	return res
}

//左子树 ---> 根结点 ---> 右子树
//BFS
func InOrder(root *TreeNode,res *[]int, k1 int, k2 int) {
	if root != nil {
		InOrder(root.Left,res,k1,k2)
		if root.Val >=k1 && root.Val <= k2 {
			*res = append(*res,root.Val)
		}
		if root.Val > k2 { //可行性剪枝
			return
		}
		InOrder(root.Right,res,k1,k2)
	}
}
