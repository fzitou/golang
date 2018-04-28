#### golang 字符串

- rune相当于go的char

```bash
package main

import "fmt"

func main(){
	s := "Yes我爱慕课网!"
	fmt.Println("字符串s长度：",len(s)) //19。中文字符，每个字符占用3字节，UTF-8编码，可变长编码
	fmt.Printf("%s\n", []byte(s)) // Yes我爱慕课网!

	fmt.Printf("%X\n", []byte(s)) // 打印出16进制
}
```

#### 字符串常见操作

- Fields,Split Join
- Contains, Index
- ToLower,ToUpper
- Trim,TrimRight,TrimLeft