package main

import (
	"errors"
	"fmt"
	"strconv"
)

const MaxTop = 20

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

func (s *Stack) Push(val int) (err error) {
	if s.IsFull() {
		return errors.New("Stack Full")
	}
	s.Top++
	s.arr[s.Top] = val
	fmt.Println("stack push ", val)
	return
}

func (s *Stack) Pop() (val int, err error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack Empty")
	}
	val = s.arr[s.Top]
	s.Top--
	return val, nil
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
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}

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
	}
	return res
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
	numStack := &Stack{
		Top: -1,
	}
	oprStack := &Stack{
		Top: -1,
	}

	exp := "30+30*6-4-6"

	// 运算需要构造
	index, a, b, opr, res := 0, 0, 0, 0, 0
	keepNum := ""
	for {

		ch := exp[index : index+1] // "3" 单个字符串, "+" ==> 43
		temp := int([]byte(ch)[0]) //字符串转为byte,  字符转的ASCII码
		if oprStack.IsOpr(temp) {  //符号
			if oprStack.IsEmpty() {
				oprStack.Push(temp)
			} else {
				if oprStack.Nice(oprStack.arr[oprStack.Top]) >=
					oprStack.Nice(temp) {
					a, _ = numStack.Pop()
					b, _ = numStack.Pop()
					opr, _ = oprStack.Pop()
					res = oprStack.Cal(a, b, opr)
					//运算结果重新入数栈
					numStack.Push(res)
					//当前符号栈压入符号栈
					oprStack.Push(temp)
				} else {
					oprStack.Push(temp)
				}
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
