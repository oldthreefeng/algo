/*
@Time : 2019/9/12 1:05
@Author : louis
@File : testep
@Software: GoLand
*/

package main

import "fmt"

func main() {
	//const (
	//	bases   float64 = 0.049            //商业贷款基准年利率
	//	baseG   float64 = 0.032            //公积金贷款基准年利率
	//	Month           = 360              //贷款月数12*30
	//	BorrowG         = 150000           // 公积金贷款总额
	//	Borrows         = 880000           //商业贷款总额
	//	bs      float64 = Borrows / Month  //商贷每月偿还本金
	//	bg      float64 = BorrowG / Month  //公积金每月偿还本金
	//	S               = bases * 1.1 / 12 //商业贷款月利率
	//	G               = baseG / 12       //公积金贷款月利率
	//)
	//var (
	//	moneyG   = [Month]float64{} //公积金每月待还
	//	moneys   = [Month]float64{} //商业每月待还
	//	money    = [Month]float64{} //总待还
	//	payBackG float64            //公积金已偿还
	//	payBacks float64            //商贷已偿还
	//	payBack  float64            //总偿还
	//)
	//for i := 0; i < Month; i++ {
	//	moneyG[i] = bg + (BorrowG-bg*(float64(i)))*G
	//	payBackG += moneyG[i]
	//	moneys[i] = bs + (Borrows-bs*(float64(i)))*S
	//	payBacks += moneys[i]
	//	money[i] = moneyG[i] + moneys[i]
	//	payBack = payBacks + payBackG
	//	fmt.Printf("第%3d月还款%.2f,\t其中公积金%.2f,\t商贷%.2f,\t已还%.2f\n",
	//		i+1, money[i], moneyG[i], moneys[i], payBack)
	//}

	fmt.Println(trailingZerosX(110))
}

//超时
func trailingZeros(n int64) int64 {
	// write your code here, try to do it without arithmetic operators.
	nFact := fact(n)
	return counterFactor(nFact,10)
}

func counterFactor(m, factor int64) (res int64) {
	for m%factor == 0 {
		res++
	}
	return res
}

func fact(n int64) (res int64) {
	res = 1
	var i int64
	for i = 1; i <= n; i++ {
		res *= 1
	}
	return res
}

/*
原理:
$$
f(n)=[\farc{n!}{5}] + [\farc{n!}{5^2}] +\cdots + [\farc{n!}{5^n}]
$$
$$
q_0 = n
q_{i+1} = [\farc{q_i}{5}]
$$
$$
q_{i+1} = 0  <==> 5^{k+1} > n
$$

for q != 0 {
	q /= 5   // ==> a1 + a2 + ... + an
	res += q
}

*/
func trailingZerosX(n int64) int64 {
	// write your code here, try to do it without arithmetic operators.
	var res int64

	for n > 0 {
		n /= 5
		res +=n
	}
	return res
}