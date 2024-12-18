## 数据类型
- 布尔 (bool)
- 数字 (int/float32/float64)
- 字符串 (string)
- 派生类型
  - 指针类型（Pointer）
  - 数组类型
  - 结构化类型(struct)
  - Channel 类型
  - 函数类型
  - 切片类型
  - 接口类型（interface）
  - Map 类型
### 管道

管道是go在语言层面提供的协程间的通信方式FIFO；每个管道只能存储一种类型的数据。

#### 声明和初始化

- 变量声明（```var ch chan int ```）值为nil

- 内置函数make()

  ```go
  ch1 := make(chan int) // 无缓冲区
  ch2 := make(chan int, 10) // 有缓冲区
  ```

#### 知识点

- 对于值为nil的管道，无论读写都会阻塞，而且是永久阻塞
- 关闭的管道可以读；写会触发panic
- 协程读取管道阻塞的条件有：管道无缓冲区、管道的缓冲区中无数据、管道的值为nil
- 协程写入管道阻塞的条件有：管道无缓冲区、管道的缓冲区已满、管道的值为nil
- 触发panic的操作：关闭值为nil的管道、关闭已经被关闭的管道、向已经关闭的管道写入数据
- 单向管道只是对管道的一种使用限制，数据结构上并没有单向管道。

#### 源码

环形队列实现：runtime/chan.go:hchan

- 关闭管道时会把recvq中的协程全部唤醒，这些协程获取的数据都为nil。同时会把sendq队列中的协程全部唤醒，这些协程会触发panic

### slice

切片又称动态数据，依托数据实现，可以方便地进行扩容和传递。

#### 声明和初始化

- 变量声明（```var s []int```）值为nil 不需要分配内存

- 字面量

  ```go
  s1 := []int{} // 空切片长度为0 需要分配内存  
  s2 := []int{1, 2, 3} // 长度和容量为3的切片
  ```

- 内置函数make()

  ```go
  s1 := make([]int, 10) // 指定长度 容量默认和长度一致
  s2 := make([]int, 10, 100) // 指定长度和容量
  ```

- 从切片和数据中切取

  ```go
  array := [5]int{1, 2, 3, 4, 5}
  s1 := array[0:2] // 与原数据共用一部分内存 容量为5
  s2 := s1[0:1] // 与原数据共用一部分内存 容量为5
  ```

#### 知识点

- 切片基于数据和切片创建，与原数组和切片共享底层空间，修改切片会影响原数组或切片
- s1 := append(s, 1) 中s可以为nil
- 当切片空间不足时，append()会先创建新的大容量切片，添加元素后再返回新切片
- 容量小于1024扩容为原来的2倍；大于或等于1024则扩大为原来的1.25倍。在该规则的基础上还会考虑元素类型与内存分配规则，对实际扩张值做一些微调
- copy()拷贝两个切片时，拷贝数量取两个切片长度的最小值。拷贝过程不会发生扩容
- 通过函数传递切片时，不会拷贝整个切片；切片本身只是一个结构体
- 谨慎使用多个切片操作同一个数组，以防读写冲突
- 在切取表达式a[low:high]中，如果a是字符串和数组则hign <= len(a) 如果a是切片则hign <= cap(a) low的默认值是0；high的默认值是a的长度
- 扩展表达式a[low:high:max]可以规避append()新切片覆盖原有切片的风险

#### 源码

数组实现：runtime/slice.go:slice

### map

#### 声明和初始化

- 变量声明（```var FruitColor map[sting]string```）值为nil

- 字面量

  ```go
  m := map[string]int {
    "apple": 2,
    "banana": 3,
  }
  ```

- 函数make() 可以指定map的容量 有效的减少内存分配次数

  ```go
  m := make(map[string]int, 10)
  m["apple"] = 2
  m["banana"] = 3
  ```

#### 知识点

- 在map为nil或指定键不存在的情况下，delete()也不会报错
- 如果键不存在，返回值为相应类的零值
- map操作不是原子的（sync.Map）
- 未初始化map添加元素会触发panic

#### 源码

hash表runtime/map.go:hmap

- 数组+链地址法（bucket可以存储8个键值对）

- 负载因子 = 键数量 / bucket数量 达到6.5会触发rehash （渐进式rehash 每次搬迁2个键值对）

- 等量扩容 overflow的bucket中大部分是空的 访问效率差

### string

#### 知识点

- 空字符串只是长度是0，但不是nil
- 单行语句拼接多个字符串只分配一次内存
- 无论是字符串转换成[]byte 还是[]byte转换成string 都发生一次内存拷贝
- utf-8编码 len()返回的是字节数
- 值不可修改

#### 源码

runtime/string.go:stringStruct

- byte切片

### select

select是go在语言层面提供的多路I/O复用机制，可以监控多个管道，当其中某一个管道可操作时就触发相应的case分支。

#### 知识点

- select{}会永久阻塞 测试会报错
- 当select的多个case语句中的管道均阻塞时，整个select语句也会陷入阻塞（没有default语句的情况），直到任意一个管道解除阻塞

#### 源码

runtime/scase

- 类型为caseNil的case语句表示其操作的管道值为nil，由于nil管道既不可读也不可写，意味着这类case永远不会命中，在运行时会被忽略，所以在case语句中值为nil的管道中写数据不会触发panic的原因。
- 如果所有的case都未命中且没有default，selectgo()将阻塞所有等待所有管道，任一管道就绪后，都将开始新的循环。

### for-range

for-range 可作用于数组、数组指针、切片、string、map和channel类型

#### 知识点

- for-range表达式在遍历开始前就已经决定了循环次数（数组、切片、string）
- for-range表达式遍历值为nil的channel是会永久阻塞；其他值为nil的类型直接结束
- range作用于数组时，从下标0开始一次遍历数组元素，返回元素的下标和元素值。如果只有一个循环变量，那么range仅返回元素的下标。作用于数组指针，效果一样
- range作用于切片时，与数组类似，返回元素的下标和元素值
- range作用于string时，仍然返回元素的下标和元素值，但由于string底层使用Unicode编码存储字符，字符可能占用1~4个字节，所以下标可能是不连续的，并且元素值是该字符对应的Unicode编码的首个字节的值
- range作用于map时，返回每个元素的key和value。不要在循环中修改map
- range作用于channel，会返回channel中所有的元素
- 当遍历数组、切片、string和map时，忽略第二个返回值，使用下标访问元素可以一定程度上提升性能

## 杂记

### 时间
格式化 "2006-01-02 15:04:05"
### 交叉编译
```cgo
export GOOS=linux
export GOARCH=amd64

unset GOOS
unset GOARCH
```
## Web框架

### go mod
- go env -w GO111MODULE=on
- go env -w GOPROXY=https://goproxy.cn,direct
### Gin
- go mod init src
- go mod tidy
- go get -u github.com/gin-gonic/gin