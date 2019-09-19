## DFS(depth first search)

```go
从起点出发
```

- 类程序

```go
func main(){
	所有点为新点
	起点为1
	终点为8
	fmt.Println(Dfs(起点))
}

// 从v出发,能否走到终点.
func Dfs(v) bool {
	if v == 终点 {
		return true
	}
	if v == 走过的点 {
		return false
	}
	将v标记为走过的点
	对v相邻的节点u,进行查询 {
		if Dfs(u) == true {
			return true
		}
	}
	return false
}
```
- 在图上寻找路径

```go
// 记录路径
type Node int
var (
	path  [MAX_LEN]Node
	depth int
)

func Dfs(v) bool {
	if v == 终点 {
		path[depth] = v
		return true
	}
	if v == 走过的点 {
		return false
	}
	将v标记为走过的点
	path[depth]=v
	depth++
	对v相邻的节点u,进行查询 {
		if Dfs(u) == true {
			return true
		}
	}
	// 将v走不到终点的路径去掉.
	depth--
	return false
}

func main(){
    depth = 0
    if Dfs(起点) {
    	for i := 0; i <= depth; i++ {
    		fmt.Print(path[i])
    	}
    }
}
```

- 遍历图上所有节点

```go
func Dfs(v) {
	if v == 走过的点 {
		return
	}
	将v标记为走过的点
	对v相邻的节点u,进行查询 {
        Dfs(u) 
    }
}

func main()  {
    所有点为标记新点
    for 图中有新点k  {
    	Dfs(k)
    }
}

```

### 图的表示方法 -- 邻接矩阵
用一个二维数组G存放图,G[i][j],表示节点i和节点j的情况:
有无边, 边方向, 权值大小等.

遍历的复杂度: O(n^2)  n为节点数目

### 图的表示方法 -- 邻接表
每一个节点V对应的一个一维数组(vector),里面存放从V连出去的边.
边的信息包括另一节点,还可能包含权值等

遍历的复杂度: O(n+e) n为节点数目, e为边数目

## 城堡问题