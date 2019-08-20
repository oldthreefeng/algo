/**
 * @time: 2019-08-19 22:37
 * @author: louis
 */
package search

import "fmt"

/*
 人有体力、情商、智商的高峰日子，它们分别每隔 23天、28天和33天出现一次。
对于每个人，我们想 知道何时三个高峰落在同一天。给定三个高峰出现的日子p,e和i（不一定是第一次高峰出现的日子）,
再给定另一个指定的日子d，你的任务是输出日子d之后，下一次三个高峰落在同一天的日子（用距离d 的天数表示）。
例如：给定日子为10，下次出现三个高峰同一天的日子是12，则输出2。

p e i 同时出现的日子;  体力高峰p 23 ; 情商高峰e 28 ; 智力高峰i 33
*/


func Period(p, e, i, d int) {
	var k int
	for k = d+1 ; k < 100000; k++ {
		if (k-i)%33 == 0 && (k-p)%22 ==0 && (k-e)%28 ==0 {
			fmt.Printf("the next triple peak occurs in %d days.", k-d)
			break
		}
	}
}