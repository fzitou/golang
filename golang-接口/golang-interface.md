#### golang interface定义

​	在Go中，**接口是一组方法签名**。 当一个类型为接口中的所有方法提供定义时，它被称为实现该接口。它与面向对象编程非常相似。接口指定类型应具有的方法，类型决定如何实现这些方法。

#### 创建和实现interface

​	在Golang中只要实现了接口定义的方法，就是实现了该interface

```go
package main

import "fmt"

// 定义接口:
type VowelsFinder interface {
	// 查看26个字母中的元音字母
	FindVowels() []rune
}

type MyString string

// 实现接口：
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}
func main() {
	name := MyString("Sam Anderson") //类型转换
	var v VowelsFinder               //定义一个接口类型的变量
	v = name
	fmt.Printf("Vowels are %c", v.FindVowels())
}

// output:
Vowels are [a e o]
Process finished with exit code 0
```



