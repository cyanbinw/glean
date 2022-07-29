package genericityGlean

import (
	"fmt"
	"reflect"
)

type GenericityDemo struct {
}

func (data *GenericityDemo) Run() error {
	useGenericity[int](1)
	fmt.Println(" -------------------- Split line -------------------- ")
	useGenericity[float32](1.1)
	fmt.Println(" -------------------- Split line -------------------- ")
	useGenericity[string]("1")
	fmt.Println(" -------------------- Split line -------------------- ")
	useGenericity[byte]('1')
	fmt.Println(" -------------------- Split line -------------------- ")
	useGenericity[bool](true)
	fmt.Println(" -------------------- Split line -------------------- ")
	useGenericity[GenericityDemo](data)
	fmt.Println(" -------------------- Split line -------------------- ")
	return nil
}

func (data *GenericityDemo) Stop() error {
	return nil
}

func (data *GenericityDemo) Close() error {
	return nil
}

func useGenericity[T any](i any) {
	fmt.Println(reflect.TypeOf(i))
	fmt.Printf("%T \n", i)
	fmt.Println(i)
}
