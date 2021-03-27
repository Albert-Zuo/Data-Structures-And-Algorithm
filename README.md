

# Data-Structures-And-Algorithm

![img](https://img.shields.io/badge/language-go-blue.svg) ![](https://img.shields.io/github/issues-raw/donng/Play-with-Data-Structures)

## Introduction

本项目基于 go 语言实现常用的数据结构，每个数据结构都有单元测试集，特别适合新手学习数据结构。后续也将持续实现一些常用的算法。

感谢 **liuyubobobo** 老师在慕课网上推出的 Java 版实战课程[《玩转数据结构》。](https://coding.imooc.com/class/207.html) Java 版本的实现请参考老师的 Github 仓库地址: https://github.com/liuyubobobo/Play-with-Data-Structures 。

## Quick Start

项目使用了 go modules 特性，因此需要确保您的 go version 在 1.11 之上。

### Download & Build

- Download: 

  -  `git clone https://github.com/Albert-Zuo/Data-Structures-And-Algorithm.git`

- Build: 

  - `go get github.com/Albert-Zuo/Data-Structures-And-Algorithm `

- Golang Unit Test: 

    - 在当前项目下，打开cmd命令行窗口
    - 进入需要测试的数据结构模块下，例如：`cd  dataStructures/array`
    - 输入命令：`go test -v`

### catalog

1、基本数据结构


|**一、线性表** | [Go源码 ](structures/array) |
| :--- | :---: |
| 1-1    定义列表接口 | [list.go](structures/array/list.go) |
| 1-2    动态数组的实现 | [arrays.go](structures/array/arrays.go) |
| 1-3-1 数组迭代器的实现 | [iterator.go](structures/array/iterator.go)|
| 1-3-2 动态数组的使用与测试 | [arrays_test.go](structures/array/arrays_test.go) |
| 1-4-1 基于数组的栈的实现 | [arraystack.go](structures/stack/arraystack.go) |
| 1-4-2 栈的测试 | [stack_test.go](structures/stack/stack_test.go) |
| 1-5     队列的实现 | [arrayqueue.go](structures/queue/array.go) |
| **二、链表** | [Go源码 ](structures/linked) |
| 2-1-1 普通链表的实现 | [linkedlist.go](structures/linked/linkedlist.go) |
| 2-1-2 普通链表的测试 | [linkedlist_test.go](structures/linked/linkedlist_test.go) |
| 2-2     基于链表的栈实现 | [linkedstack.go](structures/stack/linkedstack.go) |
| 2-3     基于链表的队列实现 | [linkedqueue.go](structures/queue/linked.go) |
| 2-4     跳表的实现 | [linkedstack.go](structures/spiklist/spiklist.go) |
| 2-5     lru cache 链表 | [lrucache.go](structures/cache/lrucache.go) |
|  |  |



2、 基本算法

| 一、排序算法         |         [Go源码 ](algorithm/sort)         |
| -------------------- | :---------------------------------------: |
| 1-1 冒泡排序         |   [bubble.go](algorithm/sort/bubble.go)   |
| 1-2 插入排序         |   [insert.go](algorithm/sort/insert.go)   |
| 1-3 快速排序         |    [quick.go](algorithm/sort/quick.go)    |
| 1-4 希尔排序         |    [shell.go](algorithm/sort/shell.go)    |
| **二、负载均衡算法** |       [Go源码 ](algorithm/balance)        |
| 2-1 轮询算法         |  [round.go ](algorithm/balance/round.go)  |
| 2-2 随机算法         | [random.go ](algorithm/balance/random.go) |
| 2-3 哈希算法         |   [hash.go](algorithm/balance/hash.go)    |
|                      |                                           |

