package main

import (
"fmt"
"strings"
)

func main() {
    s1 := "i'm ok"
    fmt.Println(s1)

    s2 := `
    世情薄
     人情恶
      雨送黄昏花易落
 `
    fmt.Println(s2)
    // 长度
    fmt.Println(len(s2))
    // 字符串拼接
    name := "zhangsan"
    world := "db"

    ss := name + world
    fmt.Println(ss)

    ss1 := fmt.Sprintf("%s%s", name, world)

    fmt.Printf("%s%s\n", name, world)
    fmt.Println(ss1)

    // 分割
    s3 := "D:\\test"
    ret := strings.Split(s3, "\\")
    fmt.Println(ret)

    // 包含
    fmt.Println(strings.Contains(ss, "lisi"))
    fmt.Println(strings.Contains(ss, "zhangs"))

    // 前缀
    fmt.Println(strings.HasPrefix(ss, "zhang"))
    // 后缀
    fmt.Println(strings.HasSuffix(ss, "ndb"))

    // 
    s4 := "abcdef"
    fmt.Println(strings.Index(s4, "c"))
    fmt.Println(strings.LastIndex(s4, "b"))

    // 拼接
    fmt.Println(strings.Join(ret, "+"))
}
