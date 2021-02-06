

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


|**一、线性表** | [Go源码 ](dataStructures/array) |
| :--- | :---: |
| 1-1 定义列表接口 | [list.go](dataStructures/array/list.go) |
| 1-2 动态数组的实现 | [arrays.go](dataStructures/array/arrays.go) |
| 1-3 数组迭代器的实现 | [iterator.go](dataStructures/array/iterator.go)|
| 1-4 动态数组的使用与测试 | [arrays_test.go](dataStructures/array/arrays_test.go) |

