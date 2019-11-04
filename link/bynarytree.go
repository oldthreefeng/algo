package main

import "fmt"

type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

//PreOrder 先输出root节点,然后再输出左子树,然后再输出右子树
func PreOrder(h *Hero) {
	if h != nil {
		fmt.Printf("no = %d name = %s \n", h.No, h.Name)
		PreOrder(h.Left) //递归
		PreOrder(h.Right)

	}

}

func PostOrder(h *Hero)  {
	if h != nil {
		PostOrder(h.Left)
		PostOrder(h.Right)
		fmt.Printf("no = %d name = %s \n", h.No, h.Name)
	}
}

func main() {
	root := &Hero{
		No:   1,
		Name: "宋江",
	}

	left1 := &Hero{
		No:   2,
		Name: "无用",
	}
	right1 := &Hero{
		No:   3,
		Name: "卢俊义",
	}
	root.Left = left1
	root.Right = right1
	right2 := &Hero{
		No:   4,
		Name: "林冲",
	}
	right1.Right = right2
	left2 := &Hero{
		Name:"case",
		No: 5,
	}
	left1.Left = left2
	PreOrder(root)
	PostOrder(root)
}

//    1
//  2   3
//5       4

// preOrder   1 2 5 3 4
// PostOrder  5 2 4 3 1
// IndexOrder 5 2 1 3 4
