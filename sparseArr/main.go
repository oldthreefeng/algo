package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ValNode struct {
	row int
	col int
	val int
}

func readFromFile(filePath string) (err error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	read := bufio.NewReader(file)
	index := 0
	var arr [][]int
	for {
		//ReadBytes直到读取输入中第一次出现delim，
		//返回包含数据的切片，包括分隔符。
		//如果ReadBytes在找到分隔符之前遇到错误，
		//它返回错误之前读取的数据和错误本身（通常是io.EOF）。
		//当且仅当返回的数据没有结束时，ReadBytes返回err！= nil
		// For simple uses, a Scanner may be more convenient.
		line, err := read.ReadBytes('\n') //读取一行
		if err != nil {
			//nil不为空说明文件读取完毕,结束读取
			break
		}
		//返回的是一个切片
		temp := strings.Split(string(line), "\t")
		// temp[0]对应的是row ,temp[1]对应的是col,temp[2]对应的是val
		// temp是一个string类型的切片
		row, _ := strconv.Atoi(temp[0])
		col, _ := strconv.Atoi(temp[1])
		val, _ := strconv.Atoi(temp[2]) // index为0的时候全是0

		fmt.Printf("%d row in file: %d\t%d\t%d\t\n", index, row, col, val)
		// 初始化arr
		index++
		if index == 1 {
			// 第一行第一个存的是行,第二个存的是列
			// 将每一行的初始化到arr切片
			for i := 0; i < row; i++ {
				var tempArr []int
				for j := 0; j < col; j++ {
					// 切片的存入数据必须用append
					tempArr = append(tempArr, val)
				}
				// 将这个临时的一维切片tempArr存入二维切片arr
				arr = append(arr, tempArr)
			}
		}

		if index != 1 {
			arr[row][col] = val
		}
	}

	fmt.Println("recover from file:")
	for _, v := range arr {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
	return
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
		//这里存入的格式对应后面分割符的格式
		_, err = writer.WriteString(fmt.Sprintf("%d\t%d\t%d\t\n", valNode.row, valNode.col, valNode.val))
	}
	_ = writer.Flush()

	//4. 将存盘文件恢复成原有数组
	// 4.1打开文件,恢复原始数组
	readFromFile(filepath)

	// 4.2使用稀疏数组恢复
	//var chessMap2 [11][11]int
	//
	//for i, valNode := range sparseArr {
	//	if i == 0 {
	//		continue
	//	}
	//	chessMap2[valNode.row][valNode.col] = valNode.val
	//}
	////5. 取出数据
	//fmt.Println("从稀疏数组恢复的数据:")
	//for _, v := range chessMap2 {
	//	for _, v2 := range v {
	//		fmt.Printf("%d\t", v2)
	//	}
	//	fmt.Println()
	//}
}
