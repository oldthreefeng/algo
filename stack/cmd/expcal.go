package main

import (
	"fmt"
	"strconv"
)

const MaxTop = 30

type Stack struct {
	Top int
	arr [MaxTop]int
}

func (s *Stack) IsEmpty() bool {
	return s.Top == -1
}

func (s *Stack) IsFull() bool {
	return s.Top == MaxTop-1
}

func (s *Stack) Size() int  {
	return len(s.arr[:s.Top])
}

func (s *Stack) Push(val int) bool {
	if s.IsFull() {
		return false
	}
	s.Top++
	s.arr[s.Top] = val
	//fmt.Printf("stack push  %#v\n", val)
	return true
}

func (s *Stack) Pop() (val int, bool bool) {
	if s.IsEmpty() {
		return 0, false
	}
	val = s.arr[s.Top]
	s.Top--
	//fmt.Printf("stach pop %#v\n", val)
	return val, true
}

func (s *Stack) List() {
	if s.IsEmpty() {
		return
	}
	//for k,v :=range s.arr {
	//	defer fmt.Printf("arr[%d] = %d\n",k,v)
	//}
	fmt.Println("Stack -->")
	for i := s.Top; i >= 0; i-- {
		fmt.Printf("arr[%d] = %d\n", i, s.arr[i])
	}
}

// 判断是否为运算符[+,-,*,/]
func (s *Stack) IsOpr(val int) bool {
	/* ASCII '42 * 43 + 45 - 47 /' */
	if val == 42 || val == 43 || val == 45 || val == 47 || val == 94 {
		return true
	} else {
		return false
	}

}

func Power(a, n int) int {
	res := 1
	for n != 0 {
		res *= a
		n--
	}
	return res
}

func (s *Stack) Cal(a, b, opr int) int {
	res := 0
	switch opr {
	case 42:
		res = b * a
	case 43:
		res = b + a
	case 45:
		res = b - a
	case 47:
		res = b / a
	case 94:
		res = Power(b, a) // b^a
	default:
		fmt.Println("opr is err")
	}
	return res
}

// 编写方法,返回运算符的优先级[自定义]

func (s *Stack) Nice(opr int) int {
	/* * / Nice返回1
	   + - Nice返回0	*/
	res := 0
	if opr == 42 || opr == 47 {
		return 1
	} else if opr == 43 || opr == 45 {
		return 0
	} else if opr == 94 { // 94	^
		return 2
	}
	return res
}

func Exp(exp string) (res int)  {
	numStack := &Stack{
		Top: -1,
	}
	oprStack := &Stack{
		Top: -1,
	}
	index, a, b, opr, res := 0, 0, 0, 0, 0
	keepNum := ""
	for {
		ch := exp[index : index+1] // "3" 单个字符串, "+" ==> 43
		//fmt.Println(ch)
		temp := int([]byte(ch)[0]) //字符串转为byte,  字符转的ASCII码
		if oprStack.IsOpr(temp) {
			// 如果operStack是空栈,直接入栈;
			// 并将数栈也pop出两个数,进行运算,
			// 将运算的结果push到数栈,符号再入符号栈
			if oprStack.IsEmpty() {
				oprStack.Push(temp)
			} else {
				//不是空栈的话,如果栈顶的运算符优先级,
				//大于当前准备入栈的运算符优先级,先pop出栈
				//例如 栈顶运算符为 * ,准备 入栈的运算符为 + ,
				//则先出栈. 并从操作栈取出两个操作数进行运算,
				//再把结果压入操作栈==>
				//继续进行比较, 如果栈顶的运算符优先级,
				//大于当前准备入栈的运算符优先级,先pop出栈.直到栈为空.
				//执行的最多次数为运算栈的Size+1次.
				//fmt.Println(oprStack.Nice(oprStack.arr[oprStack.Top]), oprStack.Nice(temp))
				for oprStack.Nice(oprStack.arr[oprStack.Top]) >=
					oprStack.Nice(temp) {
					a, _ = numStack.Pop()
					b, _ = numStack.Pop()
					opr, _ = oprStack.Pop()
					res = oprStack.Cal(a, b, opr)
					//运算结果重新入数栈
					numStack.Push(res)
					// 弹出opr运算符之后,进行判空处理,为空就直接跳出循环
					// 直接将待入栈的运算符压入符号栈.
					if oprStack.IsEmpty() {
						break
					}
				}
				oprStack.Push(temp)
			}

		} else { //数字
			//处理多位数 keepNum string,拼接
			keepNum += ch
			if index == len(exp)-1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(val)) // 3 ==> 51(ASCII)
			} else {
				if oprStack.IsOpr(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val)) // 3 ==> 51(ASCII)
					keepNum = ""
				}
			}
		}
		// 判断index是否已经扫完
		if index+1 == len(exp) {
			break
		}
		index++
	}

	// 优先级高的已经计算完

	for { //为空就弹出
		if oprStack.IsEmpty() {
			break
		}
		a, _ = numStack.Pop()
		b, _ = numStack.Pop()
		opr, _ = oprStack.Pop()
		res = oprStack.Cal(a, b, opr)
		//运算结果重新入数栈
		numStack.Push(res)
	}

	res, _ = numStack.Pop()
	return res
	//fmt.Printf("exp %s = %v", exp, res)
}

func main() {
	//1. 实现栈
	//stack := &Stack{
	//	Top: -1,
	//}
	//
	//stack.Push(1)
	//stack.Push(2)
	//stack.Push(3)
	//stack.Push(4)
	//stack.Push(5)
	//
	//stack.List()
	//val, _ := stack.Pop()
	//fmt.Println("stack pop ", val)
	//val, _ = stack.Pop()
	//fmt.Println("stack pop ", val)
	//val, _ = stack.Pop()
	//fmt.Println("stack pop ", val)
	//val, _ = stack.Pop()
	//fmt.Println("stack pop ", val)
	//val, _ = stack.Pop()
	//fmt.Println("stack pop ", val)
	//stack.List()

	//2. 加减乘除
	// 1. 实现加减乘除栈

	exp := "30+30*6-4^2-6"
	res := Exp(exp)
	fmt.Printf("exp %s = %v", exp, res)
}

//stack push  30
//stack push  43
//stack push  30
//stack push  42
//stack push  6
//stack push  180
//stack push  45
//stack push  4
//stack push  176
//stack push  45
//stack push  6
//stack push  170
//stack push  200
//exp 30+30*6-4-6 = 200
