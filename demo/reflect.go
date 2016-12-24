package main

import (
	"fmt"
	"reflect"
)

func init() {
	x := 1

	// func TypeOf(i interface{}) Type
	fmt.Println("Type:", reflect.TypeOf(x))

	fmt.Println("Value:", reflect.ValueOf(x))
}

func main() {

}
