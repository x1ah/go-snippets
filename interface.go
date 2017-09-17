package main

import (
    "fmt"
    "strconv"
)


type Human struct {
    name string
    age int
    phone string
}

type Student struct {
    Human // 匿名字段
    school string
    loan float32
}

type Employee struct {
    Human
    company string
    money float32
}

func (h Human) SayHi() {
    fmt.Println("Hi, I'm ", h.name)
}

// 默认的 string 方法
func (h Human) String() string {
    return "string: " + strconv.Itoa(h.age) + h.name
}

func (h Human) Sing(lyrics string) {
    fmt.Println("la la da la da ", lyrics)
}

// 继承重载 SayHi 方法
func (e Employee) SayHi() {
    fmt.Println("Hi, I'm ", e.name)
}

// 这个interface 被 Human，Student，Employee 实现
type Men interface {
    SayHi()
    Sing(lyrics string)
}

func main() {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}
    fmt.Println("dd: ", mike)

    var i Men
    i = mike
    i.SayHi()
    i.Sing("hhh")

    i = tom
    i.SayHi()
    i.Sing("dd")


    x := []Men{paul, sam, tom}
    for _, value := range x {
        value.SayHi()
    }
}
