package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 汉字是UTF-8字符集 一个汉字3个字节
	str := "你好 GoLang"
	fmt.Println(len(str))

	// 循环打印字符串
	for i, v := range str {
		fmt.Printf("%d = %c\n", i, v)
	}
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%d = %c \n", i, r[i])
	}

	// 字符串和整数互转
	str = "12"
	num, _ := strconv.Atoi(str)
	fmt.Println(num)
	str = strconv.Itoa(88)
	fmt.Println(str)

	// 统计子串出现次数
	count := strings.Count("GoLang JavaLang", "Lang")
	fmt.Println(count)
	// 不区分大小写比较字符串
	fmt.Println(strings.EqualFold("GoLang", "golang"))
	// 区分大小写
	fmt.Println("GoLang" == "golang")
	// 子串第一次出现的次数
	fmt.Println(strings.Index("GoLangJava", "a"))

	// 字符串替换
	fmt.Println(strings.Replace("GoLangJavaLang", "Lang", "语言", -1))
	fmt.Println(strings.Replace("GoLangJavaLang", "Lang", "语言", 1))
	// 字符串拆分
	fmt.Println(strings.Split("Go-Java", "-"))

	now := time.Now()
	fmt.Println(now)
	fmt.Printf("年 %v\n", now.Year())
	fmt.Printf("月 %v\n", now.Month())
	fmt.Printf("月 %v\n", int(now.Month()))
	fmt.Printf("日 %v\n", now.Day())
	// 各个数字必须是固定的 必须这样写
	fmt.Println(now.Format("2006-01-02 15:04:05"))

}
