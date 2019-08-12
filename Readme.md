# algo

## 稀疏数组

> 当一个数组中大部分元素为０，或者为同一个值的数组时，可以使用稀疏数组来保存该数组。
> 把具有不同值的元素的行列及值记录在一个小规模的数组中，从而 缩小程序的规模

```cgo
1) 使用稀疏数组，来保留类似前面的二维数组(棋盘、地图等等)
2) 把稀疏数组存盘，并且可以从新恢复原来的二维数组数
```
[sparseArr](sparseArr/main.go)

## 队列(Queue)

> 队列是一个有序列表，可以用 数组或是 链表来实现。
>
> 遵循 先入先出的原则。即：先存入队列的数据，要先取出。后存入的要后取出

- singleQueue

> 数组实现有序列表,需要两个变量来进行处理,front和rear,初始值都为-1
>
> front指向队首,不包含队首元素;rear指向队尾,包含队尾元素

[SingleQueue](queue/singleQueue.go)

- circleQueue

```cgo
1) 什么时候表示队列满 (tail + 1) % maxSize = head
2) tail == head 表示空
3) 初始化时， tail = 0 head = 0
4) 怎么统计该队列有多少个元素 (tail + maxSize - head ) % maxSize
5) 取出和存入数据时,head和tail该如何变化
head = (head + 1 + maxSize) % maxSize
tail = (tail + 1 + maxSize) % maxSize

```

[circleQueue](queue/circleQueue.go)


## 链表

> 链表是有序的列表
>
>一般来说，为了比较好的对单链表进行增删改查的操作，我们都会给他设置一个头结点, 头
 结点的作用主要是用来标识链表头， 本身这个结点不存放数据
 
 ```cgo
 单向链表
 
 head ==> data / &next ==> data /&next ==>
[0xc000050360,0,head,,0xc000050390] ==>
[0xc000050390,1,宋江,及时雨,0xc0000503f0] ==>
[0xc0000503f0,4,林冲,豹子头,0x0] ==>

双向链表

head <==> &pre /data / &next <==> &pre /data / &next
每个字段的含义
temp.no, temp.name, temp.nickName, temp.pre,temp, temp.next
[0,head,,0x0,0xc0000500c0,0xc000050100] <==>
[1,宋江,及时雨,0xc0000500c0,0xc000050100,0xc000050140] <==>
[2,卢俊义,玉麒麟,0xc000050100,0xc000050140,0xc000050180] <==>
[4,林冲,豹子头,0xc000050140,0xc000050180,0xc0000501c0] <==>
[3,无用,智多星,0xc000050180,0xc0000501c0,0x0] <==>
```

[singleLink](link/singlelink.go)
[doubleLink](link/doubleLink.go)