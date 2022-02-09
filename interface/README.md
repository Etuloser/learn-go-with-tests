# README

> Library reference
>
> [Blog](https://www.cnblogs.com/maji233/p/11178413.html)

## `interface`

让我们来看看这个例子： `Animal` 类型是一个接口，我们将定义一个 `Animal` 作为任何可以说话的东西。这**是 Go 类型系统的核心概念：我们根据类型可以执行的操作而不是其所能容纳的数据类型来设计抽象。**

非常简单：我们定义 `Animal` 为任何具有 `Speak` 方法的类型。`Speak` 方法没有参数，返回一个字符串。**所有定义了该方法的类型我们称它实现了 `Animal` 接口**。Go 中没有 `implements` 关键字，判断一个类型是否实现了一个接口是完全是自动地。让我们创建几个实现这个接口的类型：

```go
package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
}
func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}
func (c Cat) Speak() string {
	return "Meow!"
}

type Llama struct {
}
func (l Llama) Speak() string {
	return "?????"
}

func main() {
	animals := []Animal{&Dog{}, &Cat{}, &Llama{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
```

## `interface{}` 类型

`interface{}` 类型，**空接口**，是导致很多混淆的根源。`interface{}` 类型是没有方法的接口。由于没有 `implements` 关键字，所以所有类型都至少实现了 0 个方法，所以 **所有类型都实现了空接口**。这意味着，如果您编写一个函数以 `interface{}` 值作为参数，那么您可以为该函数提供任何值。例如：

```go
func DoSomething(v interface{}) {
   // ...
}
```

这里是让人困惑的地方：在 `DoSomething` 函数内部，`v` 的类型是什么？**新手们会认为 `v` 是任意类型的，但这是错误的。`v` 不是任意类型，它是 `interface{}` 类型。**

## 指针和接口

**Go 中的所有东西都是按值传递的。每次调用函数时，传入的数据都会被复制。对于具有值接收者的方法，在调用该方法时将复制该值**

因为所有的参数都是通过值传递的，这就可以解释为什么 `*Cat` 的方法不能被 `Cat` 类型的值调用了。任何一个 `Cat` 类型的值可能会有很多 `*Cat` 类型的指针指向它，如果我们尝试通过 `Cat` 类型的值来调用 `*Cat` 的方法，根本就不知道对应的是哪个指针。相反，如果 `Dog` 类型上有一个方法，通过 `*Dog`来调用这个方法可以确切的找到该指针对应的 `Gog` 类型的值，从而调用上面的方法。运行时，Go 会自动帮我们做这些，所以我们不需要像 C语言中那样使用类似如下的语句 `d->Speak()` 。