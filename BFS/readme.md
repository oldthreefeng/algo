

## 深度优先搜索

```cgo
相要求最优路径,要遍历所有走法.
若已经找到路径长度为n的解,则所有大于n的走法就布比尝试.

运算过程中需要存储路径上的节点,数量较少没用栈存储.

```

## 广度优先搜索

```cgo
给节点分层,起点为第0层,从起点最少需要n步,就能到达的点属于第n层.
依层次顺序,从小到大扩展. 扩展时,不能扩展已经走过的节点. 

可以确保找到最优解, 但是因扩展出的节点较多, 且多数节点都需要保存, 
因此需要存储空间较大,用队列存储节点. 

```

- 算法

```cgo
1. 把初始节点S0放入Open表中
2. 如果Open表为空,则问题无解,返回
3. 把Open表的第一个节点取出存入Closed表中,并记该节点为n
4. 考察节点n是否为目标节点, 若是,则得到问题的解,成功退出
5. 若节点n不可拓展,则转第二步
6. 扩展节点n, 将其不再Closed表和Open表中的子节点(判重)放入Open表的尾部. 
并为每一个子节点设置指向父节点的指针,然后转向第2步.

```

- 伪代码

```cgo
BFS() {
    初始化队列
    for 队列不为空且未找到目标节点 {
        取队首节点扩展,并将扩展出的非重复节点放入队尾
        必要时要记住每个节点的父节点
    }
}
```

### 抓住这头牛

```cgo
假设农夫起始位于点3，牛位于5, N=3,K=5，最右边是6。
如何搜索到一条走到5的路径？
```

### 迷宫问题

### 八数码

```cgo
把一个状态看做10进制表示
用map进行判重, 每入队一个状态, 将其加入map里面,
判重时,查找该map即可

```