/*
Copyright 2019 louis.
@Time : 2019/9/29 10:09
@Author : louis
@File : 66-preordertravelsal
@Software: GoLand

*/

package lintcode

/*
### **Description*

给出一棵二叉树，返回其节点值的前序遍历。

首个数据为根节点，后面接着是其左儿子和右儿子节点值，"#"表示不存在该子节点。节点数量不超过20


### **Example**

**样例 1:**

```
输入：{1,2,3}
输出：[1,2,3]
解释：
   1
  / \
 2   3
它将被序列化为{1,2,3}
前序遍历
```

**样例 2:**

```
输入：{1,#,2,3}
输出：[1,2,3]
解释：
1
 \
  2
 /
3
它将被序列化为{1,#,2,3}
前序遍历
```

### **Challenge**

你能使用非递归实现么？
*/

func PreOrderTravelSal(root *TreeNode) []int {
	//var s stack.Stack
	var res = make([]int,0)
	PreOrder(root,&res)
	return res
}

func PreOrder(root *TreeNode, res *[]int) {
	if root!=nil {
		*res = append(*res, root.Val)
		PreOrder(root.Left,res)
		PreOrder(root.Right,res)
	}
}

func NewTravelSal() *TreeNode {
	var t = &TreeNode{}
	t.Val = 1
	t.Left = &TreeNode{Val:2}
	t.Left.Left = &TreeNode{Val:3}
	t.Left.Right = &TreeNode{Val:4}
	t.Right = &TreeNode{Val:5}
	t.Right.Left = &TreeNode{Val:6}
	t.Right.Right = &TreeNode{Val:7}
	return t
}