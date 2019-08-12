package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	//1. 创建原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑色
	chessMap[2][3] = 2 //白色
	//fmt.Println("originArr:")
	//for _, v := range chessMap {
	//	for _, v2 := range v {
	//		fmt.Printf("%d\t", v2)
	//	}
	//	fmt.Println()
	//}
	//2. 转成稀疏数组. 想-> 算法
	// 遍历chessMap,发现有不为0的元素,将其放入到对应的切片中
	// 标准的稀疏数组还有一个记录原始二维数组的规模(行和列,默认值)
	var sparseArr []ValNode
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				valNode = ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	fmt.Println("sparseArr:")
	for i, valNode := range sparseArr {
		fmt.Printf("%d:\t%d\t%d\t%d", i, valNode.row, valNode.col, valNode.val)
		fmt.Println()
	}

	//3. 将稀疏数组存盘"D:\text"
	filepath := "D:/text/chessMap.data"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err =%v\n", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, valNode := range sparseArr {
		_, err = writer.WriteString(fmt.Sprintf("%d\t%d\t%d\n", valNode.row, valNode.col, valNode.val))
	}
	_ = writer.Flush()

	//4. 将存盘文件恢复成原有数组
	// 4.1打开文件,恢复原始数组
	file, err = os.OpenFile(filepath, os.O_RDONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err =%v\n", err)
	}

	read := bufio.NewReader(file)
	for {
		str, err := read.ReadString('\n')
		if err == io.EOF {
			break
		}
	}
	// 4.2使用稀疏数组恢复
	var chessMap2 [11][11]int

	for i, valNode := range sparseArr {
		if i == 0 {
			continue
		}
		chessMap2[valNode.row][valNode.col] = valNode.val
	}
	//5. 取出数据
	fmt.Println("恢复后的数据:")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
