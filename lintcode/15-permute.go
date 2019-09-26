/*
@Time : 2019/9/26 17:37
@Author : louis
@File : 15-permute
@Software: GoLand
*/

package lintcode

//This is NOT WORK
//TODO
func Permute(nums *[]int) [][]int {
	res := make([][]int,0)
	if len(*nums) <=0 {
		return res
	}
	cur := make([]int,len(*nums))
	used := make([]int,len(*nums))
	dfs(&res,&cur,nums,&used)
	return res
}

func dfs(res *[][]int, cur *[]int, nums *[]int, used *[]int) {
	if len(*cur) == len(*nums) {
		*res = append(*res,*cur)
		return
	}
	for i := 0; i < len(*nums); i++ {
		if (*used)[i] == 0 {
			*cur = append(*cur,(*nums)[i])
			(*used)[i] = 1
			dfs(res,cur,nums,used)
			(*used)[i] = 0
			*cur = pop(*cur,(*nums)[i])
		}

	}
}

func pop(a []int, n int) []int  {
	j :=0
	for _,val := range a {
		if val == n {
			a[j] =val
			j++
		}
	}
	return a[:j]
}
