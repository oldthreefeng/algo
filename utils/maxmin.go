/*
@Time : 2019/9/19 0:34
@Author : louis
@File : maxmin
@Software: GoLand
*/

package utils

func Min(nums ...int) int {
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if m > nums[i] {
			m = nums[i]
		}
	}
	return m
}

func Max(nums ...int) int {
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if m < nums[i] {
			m = nums[i]
		}
	}
	return m
}

func Power(a, b int) int {
	res := 1
	if b == 0 {
		return res
	}
	for b > 0 {
		b--
		res *= a
	}
	return res
}

func BigMulti(a, b string) string {
	if a == "0" || b == "0" {
		return "0"
	}
	// string转换成[]byte，容易取得相应位上的具体值
	bsi := []byte(a)
	bsj := []byte(b)

	temp := make([]int, len(bsi)+len(bsj))
	//两数相乘，结果位数不会超过两乘数位数和，即temp的长度只可能为 len(num1)+len(num2) 或 len(num1)+len(num2)-1
	// 选最大的，免得位数不够
	for i := 0; i < len(bsi); i++ {
		for j := 0; j < len(bsj); j++ {
			// 对应每个位上的乘积，直接累加存入 temp 中相应的位置
			temp[i+j+1] += int(bsi[i]-'0') * int(bsj[j]-'0')
		}
	}

	//统一处理进位
	for i := len(temp) - 1; i > 0; i-- {
		temp[i-1] += temp[i] / 10 //对该结果进位（进到前一位）
		temp[i] = temp[i] % 10    //对个位数保留
	}

	// a 和 b 较小的时候，temp的首位为0
	// 为避免输出结果以0开头，需要去掉temp的0首位
	if temp[0] == 0 {
		temp = temp[1:]
	}
	//转换结果：将[]int类型的temp转成[]byte类型,
	//因为在未处理进位的情况下，temp每位的结果可能超过255(go中，byte类型实为uint8，最大为255),所以temp选用[]int类型
	//但在处理完进位后，不再会出现溢出
	res := make([]byte, len(temp)) //res 存放最终结果的ASCII码

	for i := 0; i < len(temp); i++ {
		res[i] = byte(temp[i] + '0')
	}

	return string(res)
}
