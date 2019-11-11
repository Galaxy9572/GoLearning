package chapter6

import "fmt"

// 定义一个接口
// 接口由使用者定义
type Vehicle interface {
	Run()
}

type FlyingObject interface {
	Fly()
}

// 接口变量中包含两部分
// 1、实现着的类型
// 2、实现者的值（或者一个指向实现着的指针）
// interface{}表示任何类型
type Car struct {
	Brand string
}

// Car实现Run方法，接口的实现是隐式的，不需要声明实现了哪个接口
func (car Car) Run() {
	fmt.Println("Car is running...")
}

// duck typing
// 描述事物的外部行为

// 接口的组合：新建一个接口，里面放已经声明的接口类型

type FlyingVehicle interface {
	Vehicle
	FlyingObject
}
