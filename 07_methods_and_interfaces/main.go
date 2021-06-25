package main

import (
	"fmt"
)

type Struct struct {
	s string
}

func (s Struct) valueMethod() {
	fmt.Println("valueMethod():", s.s)
}

func (s *Struct) pointerMethod() {
	fmt.Println("pointerMethod():", s.s)
}

type pointerMethodInterface interface {
	pointerMethod()
}

type valueMethodInterface interface {
	valueMethod()
}

func main() {
	pointerStruct := &Struct{s: "pointerStruct"}
	valueStruct := Struct{s: "valueStruct"}
	pointerStruct.pointerMethod()
	valueStruct.valueMethod()
	pointerStruct.valueMethod()
	valueStruct.pointerMethod()

	var emptyInterfaceInstancePointerStruct interface{} = pointerStruct
}
