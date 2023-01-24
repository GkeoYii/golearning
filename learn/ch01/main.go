package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}

// 基础类型: 整型、浮点数、布尔值、字符串、字节、空值
const (
// var ch02 bool = true
// var i int = 10
// var f float32 = 1.2
// var b byte = 'a'
// var s string = "abc"
// var ok bool = true
//
// ba := []byte(s)
)

// 空值结构：if、for、switch

func ifTest() {
	i := 10
	if i > 10 {
		fmt.Println(i)
	} else {
		fmt.Println(i)
	}
}

func switchTest() {
	switch i := 6; {
	case i > 10:
		fmt.Println(i)
	case 5 < i && i < 10:
		fmt.Println(i)
		fallthrough //继续往下执行
	default:
		fmt.Println(i)
	}
}

func forTest() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// 集合类型：array、slice、map

func arrayTest() {
	//a1 := [4]int{1, 2, 3, 4}
	//var a2 [4]int = [4]int{1, 2, 3, 4}
}

func sliceTest() {
	//a1 := []int{1, 2, 3, 4}
	//// 前闭后开[1:3)
	//s1 := a1[1:3]
	//s2 := a1[:3]
	//s3 := a1[1:]
	//
	//s4 := make([]int, 4, 8)
	//s5 := []int{1, 2, 3, 4}
	//s6 := append(s2, 6, 7)
	//s7 := append(s1, s5)
}

func mapTest() {
	//m1 := make(map[string]int)
	//m2 := map[string]int{"xxx": 10}
	//m2["yy"] = 20
	//i, ok := m2["xxx"]
	//if ok {
	//	delete(m2, "xx")
	//}

}

// 函数名称首字母小写代表私有函数，只有在同一个包中才可以被调用；
// 函数名称首字母大写代表公有函数，不同的包也可以调用；
func funcTest(params ...int) {
	for _, v := range params {
		fmt.Println(v)
	}

	//闭包
	sum := func(a, b int) int {
		return a + b
	}
	fmt.Println(sum(1, 2))

	cl := closure()
	fmt.Println(cl()) // 1
	fmt.Println(cl()) // 2
	fmt.Println(cl()) // 3
}

// 每次调用i + 1
func closure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

type person struct {
	name string
	age  int
	addr address
}

type address struct {
	province string
	city     string
}

func structTest() {
	var p person
	p = person{"xxx", 11, address{
		city:     "xx",
		province: "yy"}}
	p.name = "yy"
	p.age = 10
	p = person{age: 20}
	fmt.Println(p)
	fmt.Println(p.addr.province)

	// 当值类型作为接收者时，person 类型和*person类型都实现了该接口。
	// 当指针类型作为接收者时，只有*person类型实现了该接口。
	// person String 也实现了 fmt.Stringer的String方法
	printString(&p)
	printString(p)
	//p1 := NewPerson("yy")
	//printString(&p)
}

type Stringer interface {
	String() string
}

func (p person) String() string {
	return fmt.Sprintf("the name is %s,age is %d", p.name, p.age)
}

//func (p *person) String() string {
//	return fmt.Sprintf("the name is %s,age is %d", p.name, p.age)
//}

func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

// factory
func NewPerson(name string) *person {
	return &person{name: name}
}

// 组合

type Walk interface {
	Walk()
}
type Run interface {
	Run()
}
type WalkRun interface {
	Walk()
	Run()
}

func (p person) Walk() {
	fmt.Println("walk", p.name)
}

func (p person) Run() {
	fmt.Println("run", p.name)
}

func panicTest(ip string) {
	defer func() {
		//恢复异常
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	fmt.Println("aa")
	if ip == "" {
		//异常奔溃
		panic("ip cannot empty")
	}
}
