package utils

//Gcd 欧几里德算法
//欧几里德算法又称辗转相除法，用于计算两个整数a,b的最大公约数。其计算原理依赖于下面的定理：
//
//定理：gcd(a,b) = gcd(b,a mod b)
//
//证明：a可以表示成a = kb + r，则r = a mod b
//假设d是a,b的一个公约数，则有
//d|a, d|b，而r = a - kb，因此d|r
//因此d是(b,a mod b)的公约数
//
//假设d 是(b,a mod b)的公约数，则
//d | b , d |r ，但是a = kb +r
//因此d也是(a,b)的公约数
//
//因此(a,b)和(b,a mod b)的公约数是一样的，其最大公约数也必然相等，得证
func Gcd(a, b int) int {
	if a < 0 || b < 0 {
		return -1
	}
	if a%b != 0 {
		return Gcd(b, a%b)
	}
	return b
}

func Lcm(a, b int) int {
	return a * b / Gcd(a, b)
}

func ExGcd(a, n int) int {
	var p, q = a, n
	// 17 ,26
	var x, y = 0, 1
	var z = q / p // 26/17 = 1
	for p != 1 && q != 1 {
		t := p      //17
		p = q % p   //9
		q = t       //17
		t = y       // 1
		y = x - y*z // 0-1*1
		x = t       // 1
		z = q / p   // 17/9 = 1
	}
	y %= n
	if y < 0 {
		y += n
	}
	return y
}
