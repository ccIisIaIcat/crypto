package main

import (
	"fmt"
	"reflect"
)

type myobject struct {
	Name string
	Age  int
}

func main() {
	h := myobject{"测试", 20}
	fmt.Println(h)
	hofvalue := reflect.ValueOf(h)
	ttt := reflect.TypeOf(h)
	//获取结构体的reflect.Value对象。
	for i := 0; i < hofvalue.NumField(); i++ { //循环结构体内字段的数量
		//获取结构体内索引为i的字段值
		fmt.Println(ttt.Field(i).Name)
		fmt.Println(hofvalue.Field(i))
		fmt.Println(hofvalue.Field(i).Interface().(string))
	}
	fmt.Println(hofvalue.Field(1).Type()) //获取结构体内索引为1的字段的类型

}
