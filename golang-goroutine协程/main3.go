package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
* @author wangpengcheng@ccssoft.com.cn
* @date 2018/5/16 22:50
 */

func main() {
	var a [3]int
	for i := 0; i < 10; i++ {
		// 如果下面的匿名函数不传递参数i进去则报错：panic: runtime error: index out of range
		go func() {
			a[i]++
			runtime.Gosched()
		}() // 如果下面的匿名函数不传递参数i进去则报错：panic: runtime error: index out of range
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

// output:
// panic: runtime error: index out of range
