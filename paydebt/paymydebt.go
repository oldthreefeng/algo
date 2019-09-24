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
	bases  float64 = 0.049           //商业贷款基准年利率
	baseS          = bases * 1.1     //上浮10%
	baseG          = 0.032           //公积金贷款基准年利率
	S              = baseS / 12      //商业贷款月利率上浮 10%
	G              = baseG / 12      //公积金贷款月利率
	Start          = 10 - offset     //起始还贷日期如:10月份
	offset         = 1               //偏移量.%12后为0-11.
	Year           = 30              //贷款年份
	Month          = 12*Year + Start //贷款月数12*30
	AG             = 150000          //公积金贷款总额
	AS             = 880000          //商业贷款总额
	//S               = 0.0044916 //商业贷款月利率上浮 10%
)

/*
原理:
每个月偿还本金都是一样的,利息越来越少.

总贷款金额为A,还款月数为Month,月利率为S.
故:
$$
A_n = \frac{A}{Month} + (A - \frac{A}{Month} * n) * S
$$
*/
func benJ() {
	const (
		bs = AS / (Month - Start) //商贷每月偿还本金
		bg = AG / (Month - Start) //公积金每月偿还本金
	)
	var (
		jG               = [Month]float64{} //公积金每月待还
		jS               = [Month]float64{} //商业每月待还
		money            = [Month]float64{} //总待还
		payBackG float64 = 0                //公积金已偿还
		payBacks float64 = 0                //商贷已偿还
		payBack  float64 = 0                //总偿还

	)
	for i := Start; i < Month; i++ {
		jG[i] = bg + (AG-bg*(float64(i-Start)))*G
		payBackG += jG[i]
		jS[i] = bs + (AS-bs*(float64(i-Start)))*S
		payBacks += jS[i]
		money[i] = jG[i] + jS[i]
		payBack = payBacks + payBackG
		//偏移量补足即可
		fmt.Printf("%4d年%2d月还款: %.2f = %.2f[公积金] + %.2f[商贷],已还%.2f\n",
			2019+i/12, i%12+offset, money[i], jG[i], jS[i], payBack)
	}
	fmt.Printf("等额本金共计偿还利息为:%.2f\n", payBack-AS-AG)
	fmt.Printf("等额本金公积金共计还款:%.2f\n",payBackG)
	fmt.Printf("等额本金商业贷共计还款:%.2f",payBacks)
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
func benX() {
	var es, eg float64
	es = AS * S * math.Pow(1+S, Month) / (math.Pow(1+S, Month) - 1)
	eg = AG * G * math.Pow(1+G, Month) / (math.Pow(1+G, Month) - 1)
	fmt.Printf("%.2f+%.2f = %0.2f\n", eg, es, es+eg)
	fmt.Printf("共计偿还利息为:%.2f", (es+eg)*Month-AG-AS)
}

func main() {
	benJ()
	fmt.Printf("\n等额本息:")
	benX()
}
/*
20180911132802-cd97c894b2cc5*/