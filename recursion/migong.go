package recursion

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


