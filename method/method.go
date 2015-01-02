package main

import (
    "fmt"
)


type People struct {
    name string
    age  int
}

func (p *People) Say(msg string) {
    fmt.Printf("%s:\tsay[%s]\n", p.name, msg)
}

func main() {

    taro := People{"taro", 18}
    jiro := People{name:"jiro", age: 16}
    ken_p := &People{"ken", 24}  // pointer

    fmt.Printf("jiro age:%d\n", jiro.age)
    fmt.Printf("taro age:%d\n", taro.age)
    fmt.Printf("ken age:%d\n", ken_p.age)

    taro.Say("i am taro")
}

