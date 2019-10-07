/*
Copyright 2019 louis.
@Time : 2019/9/29 23:09
@Author : louis
@File : 368-evaluateExpression
@Software: GoLand

*/

package lintcode

/*
1）如果遇到操作数，我们就直接将其输出。

2）如果遇到操作符，则我们将其放入到栈中，遇到左括号时我们也将其放入栈中。

3）如果遇到一个右括号，则将栈元素弹出，将弹出的操作符输出直到遇到左括号为止。注意，左括号只弹出并不输出。

4）如果遇到任何其他的操作符，如（“+”， “*”，“（”）等，
从栈中弹出元素直到遇到发现更低优先级的元素(或者栈为空)为止。
弹出完这些元素后，才将遇到的操作符压入到栈中。
有一点需要注意，只有在遇到" ) "的情况下我们才弹出" ( "，
其他情况我们都不会弹出" ( "。

5）如果我们读到了输入的末尾，则将栈中所有元素依次弹出。
*/
//
//func EvaluateExpress(s []string) int {
//	if s == nil || len(s) == 0 {
//		return 0
//	}
//	s = EvalRPN(s)
//	var st  stack.Stack
//	//no more '(' ')'.
//	for _,v := range  s {
//		if IsOpr(v) {
//			t1 := st.Pop().(int)
//			t2 := st.Pop().(int)
//			switch v[0] {
//			/* ASCII '42 * 43 + 45 - 47 /' */
//			case 43:
//				st.Push(t2+t1)
//			case 47:
//				st.Push(t2-t1)
//			case 42:
//				st.Push(t2*t1)
//			case 45:
//				st.Push(t2/t1)
//			}
//		} else {
//			st.Push(Atoi(v))
//		}
//	}
//	return st.Pop().(int)
//}

//func EvalRPN(s []string) []string  {
//	var str []string
//	var sa stack.Stack
//	if s == nil || len(s) == 0 {
//		return s
//	}
//	for k,v := range s {
//		if v[0] >= '0' && v[0]<='9' {
//			//数字直接进.
//			str = append(str, v)
//			continue
//		}
//		switch v[0] {
//		case '(':
//
//		}
//	}
//	return nil
//}

func IsOpr(s string) bool {
	var val byte
	val = s[0]
	if val == 42 || val == 43 || val == 45 || val == 47 || val == 94 {
		return true
	} else {
		return false
	}
}

func Nice(s int) int {
	var res = -1
	switch s {
	case '(':
		res = 0
	case '+':
		res = 1
	case '-':
		res = 1
	case '*':
		res = 2
	case '/':
		res = 2
	}
	return res
}

/*
30号, 上午/ 下午去打扫卫生
1号,  整理家务,晒晒被子
2号见面, 晚上来我家吃饭
3号去爬山, 晚上去外派家,
4号去爷爷家.
5号杭州.
*/