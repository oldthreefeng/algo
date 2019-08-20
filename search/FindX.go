package search

import "fmt"

//f(x) = x^3 -5*x^2 +10x -80 =0
//若求出的跟是a, 则要求精度 |f(a)| <= 1e-6 (10^-6)

/*

 */
func f(x float64) float64 {
	a := x*x*x - 5*x*x + 10*x - 80
	return a
}

func FindX() float64 {
	var root, x1, x2, y float64
	const EPS = 1e-6
	fmt.Println(EPS)
	x1, x2 = 0, 100
	root = x1 + (x2-x1)/2
	triedTimes := 1
	y = f(root)

	for y > EPS || -y > EPS { //精度保持在EPS范围内
		if y > 0 { // 说明
			x2 = root
		} else {
			x1 = root
		}
		root = x1 + (x2-x1)/2
		y = f(root)
		triedTimes++
	}
	fmt.Printf("root =%.8f times: %d", root, triedTimes)
	return root
}
