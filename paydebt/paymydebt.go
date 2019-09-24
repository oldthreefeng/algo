/*
@Time : 2019/9/24 16:48
@Author : louis
@File : paymydebt
@Software: GoLand
*/

package main

import (
	"fmt"
	"math"
)

const (
	bases   float64 = 0.049            //商业贷款基准年利率
	baseG   float64 = 0.032            //公积金贷款基准年利率
	Month           = 360              //贷款月数12*30
	BorrowG         = 150000           // 公积金贷款总额
	Borrows         = 880000           //商业贷款总额
	bs      float64 = Borrows / Month  //商贷每月偿还本金
	bg      float64 = BorrowG / Month  //公积金每月偿还本金
	S               = bases * 1.1 / 12 //商业贷款月利率上浮 10%
	G               = baseG / 12       //公积金贷款月利率
)

/*
原理:
每个月偿还本金都是一样的,利息越来越少.

总贷款金额为A,还款月数为Month,月利率为S.
故:
$$A_n = \frac{A}{Month} + (A - \frac{A}{Month} * n) * S$$
*/
func benJin() {
	var (
		jG               = [Month]float64{} //公积金每月待还
		js               = [Month]float64{} //商业每月待还
		money            = [Month]float64{} //总待还
		payBackG float64 = 0                //公积金已偿还
		payBacks float64 = 0                //商贷已偿还
		payBack  float64 = 0                //总偿还
	)
	for i := 0; i < Month; i++ {
		jG[i] = bg + (BorrowG-bg*(float64(i)))*G
		payBackG += jG[i]
		js[i] = bs + (Borrows-bs*(float64(i)))*S
		payBacks += js[i]
		money[i] = jG[i] + js[i]
		payBack = payBacks + payBackG
		fmt.Printf("第%3d月还款%.2f,\t其中公积金%.2f,\t商贷%.2f,\t已还%.2f\n",
			i+1, money[i], jG[i], js[i], payBack)
	}
	fmt.Printf("共计偿还利息为:%.2f\n", payBack-Borrows-BorrowG)
}

/*

原理:

每个月总偿还金额是不变的.

$A_n$为第n个月欠银行贷款. X为每个月还款金额(本金+利息).

故:

$$A_1 = A\beta(1+\beta)-X$$

$$A_2 = A_1(1+\beta)-X = A(1+\beta)^{2}-X[(1+\beta)+1]$$

$$A_n = A_{n-1}(1+\beta)-X $$

​	$$ = A(1+\beta)^n-X[(1+\beta)^{n-1}+(1+\beta)^{n-2}+\cdots+(1+\beta)+1] $$

​	$$ = A(1+\beta)^n-\frac{X[1+\beta]^n-1}{\beta}$$

因此:

$$X = \frac{A\beta(1+\beta)^{m}}{(1+\beta)^m-1}$$

*/
func benE() {
	var es, eg float64
	es = Borrows * S * math.Pow(1+S, Month) / (math.Pow(1+S, Month) - 1)
	eg = BorrowG * G * math.Pow(1+G, Month) / (math.Pow(1+G, Month) - 1)
	fmt.Printf("%.2f+%.2f = %0.2f\n", es, eg, es+eg)
	fmt.Printf("共计偿还利息为:%.2f", (es+eg)*Month-BorrowG-Borrows)
}

func main() {
	benJin()
	benE()
}
