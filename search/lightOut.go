/**
 * @time: 2019-08-19 23:15
 * @author: louis
 */
package search

/*
–第一行是一个正整数N, 表示需要解决的案例数
–每个案例由5行组成, 每一行包括6个数字
–这些数字以空格隔开, 可以是0或1
–0 表示灯的初始状态是熄灭的
–1 表示灯的初始状态是点亮的

*/

//TODO

//GetBit get bit in byte
func GetBit(c byte, i uint) byte  { //返回 byte的二进制第i位的数
	return (c >> i) & 0x1
}

func SetBit(c byte,i uint,v int) byte {
	if v != 0 {
		c |= 1<<i
		/**
		将某一位设置为1，例如设置第8位，从右向左数需要偏移7位,注意不要越界
		1<<7=1000 0000 然后与c逻辑或|,偏移后的第8位为1，逻辑|运算时候只要1个为真就为真达到置1目的
		*/
	} else {
		/**
		将某一位设置为0，例如设置第4位，从右向左数需要偏移3位,注意不要越界
		1<<3=0000 1000 然后取反得到 1111 0111 然后逻辑&c
		*/
		c = c&^(1<<i)
	}

	return c

}

func FlipBit(c byte, i uint) byte {
	c = c ^ (1 << i)
	return c
	/*
	1左移3位.
	如获取c的第4位 1<<3=0000 1000
	0000 1000 ^ 0001 1110 = 0001 0110
	*/

}

func OutputResult()  {
	
}

//- [related](https://studygolang.com/articles/14276)