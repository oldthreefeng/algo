package main

import "fmt"

// i,j进行测试
func SetWay(m *[18][17]int, i, j int) bool {

	if m[16][15] == 2 {
		return true
	} else {
		// 继续找
		if m[i][j] == 0 { //这里没有走
			// 假设这个点可以通,上下左右走
			m[i][j] = 2
			//if SetWay(m,i-1,j) {
			//	return true
			//} else if SetWay(m,i+1,j) {
			//	return true
			//} else if SetWay(m,i,j-1) {
			//	return true
			//} else if SetWay(m,i,j+1) {
			//	return true  SetWay(m,i+1,j)  SetWay(m,i-1,j)
			if SetWay(m,i+1,j) { //下右上左
				return true
			} else if SetWay(m,i,j+1) {
				return true
			} else if  SetWay(m,i-1,j) {
				return true
			} else if SetWay(m,i,j-1) {
				return true

			} else { //是死路
				m[i][j] =3
				return false
			}
		} else {
			//说明是墙
			return false
		}
	}
}

func main() {
	//二维数组模拟迷宫
	var m [18][17]int
	// 定下规则,val为1时.为墙;val为0,这条路没有走;val为2时为走过的点;如果元素的值为3;走过的但是走不通
	for i := 0; i < 17; i++ { //最上和最下置为1
		m[0][i] = 1
		m[17][i] = 1
	}
	for i := 0; i < 18; i++ { //最左和最右置为1
		m[i][0] = 1
		m[i][16] = 1
	}
	m[3][1]=1
	m[3][2]=1
//TODO 优化最佳线路
	// 堵死
	m[1][2] =1
	m[2][2] =1

	SetWay(&m,1,1)
	fmt.Println("探测完毕")

	for _, v := range m {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
	return
}
