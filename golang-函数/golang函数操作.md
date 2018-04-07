#### golang函数

 - golang函数可以返回多个值;返回值写在最后面
 - golang函数返回多个值时可以起名字，仅用于非常简单的函数，对于调用者来说没区别，随便用什么变量名称接收返回值;
 - golang主要还是函数式编程语言,是golang中的一等公民。函数的返回值，函数体内都可以是函数，函数的参数也可以是函数
 - golang中没有函数重载，操作符重载，默认参数等花哨的语法，它由一个可变参数列表的的函数语法sumArgs(values...int) int

操作示例

```go
package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 返回一个值的函数
func eval(a,b int,op string) (int,error) {
	switch op{
	case "+":
		return a+b,nil
	case "-":
		return a-b,nil
	case "*":
		return a*b,nil
	case "/":
		//return a/b
		q,_:=div(a,b)
		//只能返回一个值
		return q,nil
	default:
		return 0,fmt.Errorf("不支持运算: %s", op)
		//panic("不支持的算数运算："+op)
	}
}

// 返回2个值的函数
// 13/3=4...1
//func div(a,b int) (int,int) {
// 给返回值取名字，q是商，r是余数,建议这种直接return
func div(a,b int) (q,r int) {
	return a / b, a%b
}

// 函数返回值取名之后直接,不建议分开，建议上面一种方式
func div2(a,b int) (q, r int){
	q = a/b
	r =  a%b
	return
}

// 参数也是函数
func apply(op func(int,int) int,a,b int) int{
	// 反射
	p:=reflect.ValueOf(op).Pointer()
	opName:=runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args \n" + "(%d, %d)", opName,a,b)
	return op(a,b)
}

// 重写系统函数
func pow(a,b int) int {
	return int(math.Pow(float64(a),float64(b)))
}

// 可变参数长度函数
func sum(numbers ...int) int{
	s:=0
	for i:=range numbers{
		s+=numbers[i]
	}
	return s
}
func main(){
	fmt.Println(eval(1,2,"+"))

	//fmt.Println("13除3的结果",div(13,3)) //错误的写法，div返回2个值就不能在前面加字符串了。
	fmt.Println(div(13,3))
	// 如果函数div返回值取了名字，则调用此函数的时候，编辑器一般会自动把函数返回值赋值给其名称，自动生成
	q,r:=div(13,3)
	fmt.Println(q,r)

	fmt.Println(div2(13,3))


	if result,err:=eval(3,4,"x");err!=nil{
		fmt.Println("出错了",err)
	}else{
		fmt.Println("成功了",result)
	}

	fmt.Println(apply(pow,3,4))

	// 直接使用匿名函数
	fmt.Println(apply(
		func(a int,b int) int{
			return int(math.Pow(float64(a),float64(b)))
		},3,4))

	//
	fmt.Println(sum(1,2,3,4,5))
}
```



​	