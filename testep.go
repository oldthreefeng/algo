/*
@Time : 2019/9/12 1:05
@Author : louis
@File : testep
@Software: GoLand
*/

package main

import "fmt"

func main() {
	const (
		bases   float64 = 0.049            //商业贷款基准年利率
		baseG   float64 = 0.032            //公积金贷款基准年利率
		Month           = 360              //贷款月数12*30
		BorrowG         = 150000           // 公积金贷款总额
		Borrows         = 880000           //商业贷款总额
		bs      float64 = Borrows / Month  //商贷每月偿还本金
		bg      float64 = BorrowG / Month  //公积金每月偿还本金
		S               = bases * 1.1 / 12 //商业贷款月利率
		G               = baseG / 12       //公积金贷款月利率
	)
	var (
		moneyG           = [Month]float64{} //公积金每月待还
		moneys           = [Month]float64{} //商业每月待还
		money            = [Month]float64{} //总待还
		payBackG float64 = 0                //公积金已偿还
		payBacks float64 = 0                //商贷已偿还
		payBack  float64 = 0                //总偿还
	)
	for i := 0; i < Month; i++ {
		moneyG[i] = bg + (BorrowG-bg*(float64(i)))*G
		payBackG += moneyG[i]
		moneys[i] = bs + (Borrows-bs*(float64(i)))*S
		payBacks += moneys[i]
		money[i] = moneyG[i] + moneys[i]
		payBack = payBacks + payBackG
		fmt.Printf("第%3d月还款%.2f,\t其中公积金%.2f,\t商贷%.2f,\t已还%.2f\n",
			i+1, money[i], moneyG[i], moneys[i], payBack)
	}
}
