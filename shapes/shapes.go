/**
Library reference:
https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.5.md

方法是一个特殊函数，它的第一个参数是隐式的，被称为接收者
"A method is a function with an implicit first argument, called a receiver."

method和字段一样，从属于它的接收者

method的名字可以一样，接收者不一样就是不同的method
method可以访问接收者的字段

接收者可以是指针，这样才能正确的对接收者的实例进行操作
**/
package shapes

import "math"

// 定义了一个名为 Rectangle 的 struct
type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Radius float64
}
type Triangle struct {
	Base   float64
	Height float64
}
type Trapezoid struct {
	Top    float64
	Bottom float64
	Height float64
}
type Shape interface {
	Area() float64
}

// 定义了一个函数，它的第一个参数是隐式的 r Rectangle，即接收者为 Retangle，所以该函数是 Retangle 的一个方法，方法名为 Area
// 可以通过 . 来访问接收者 Retangle 中的字段
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 虽然同名，但接收者不一样就是不同的 method
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Triangle) Area() float64 {
	return (c.Base * c.Height) * 0.5
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
