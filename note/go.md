#### 数据类型
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

## Web框架
#### go mod
- go env -w GO111MODULE=on
- go env -w GOPROXY=https://goproxy.cn,direct
#### Gin
- go mod init src
- go mod tidy
- go get -u github.com/gin-gonic/gin