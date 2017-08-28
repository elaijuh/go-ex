package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	A string `tag_name:"tag 1"`
	B int    `tag_name:"tag 2"`
	C string `tag_name:"tag 3"`
}

func main() {

	f := Foo{
		A: "mo",
		B: 23,
		C: "po",
	}

	// https://blog.golang.org/laws-of-reflection

	// law 1 - from interface value to reflection object

	fmt.Println("type:", reflect.TypeOf(f)) // type: main.Foo
	v := reflect.ValueOf(f)
	fmt.Println("type:", v.Type())    // type: main.Foo
	fmt.Println("value:", v.String()) // value: <main.Foo Value>
	fmt.Println(v.Kind())             // struct
	fmt.Println("value:", v)          // value: {mo 23 po}

	// law 2 - from reflection object back to interface value
	fmt.Println(v.Interface()) // {mo 23 po}

	// law 3 - to modify a reflection object, the value must be settable
	vp := reflect.ValueOf(&f)
	fmt.Println("settability of vp:", vp.CanSet()) // settablility of vp: false
	vv := vp.Elem()                                // vv is the underlying element where vp point to
	fmt.Println("settability of vv:", vv.CanSet()) // settablility of vv: true

	for i := 0; i < vv.NumField(); i++ {
		f := vv.Field(i)
		ft := vv.Type().Field(i)
		tag := ft.Tag
		fmt.Printf("%d: %s %s = %v, %v\n", i, ft.Name, f.Type(), f.Interface(), tag.Get("tag_name"))
	}

	vv.Field(1).SetInt(7)
	fmt.Println("f is now:", f) // &{mo 7 po}
}
