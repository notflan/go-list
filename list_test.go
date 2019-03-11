package list

import (
	"testing"
	"fmt"
)

func TestList(t *testing.T) {
	list := New()
	list.Add(1, 2, "three", 4.4)
	fmt.Println(list)
	fmt.Println(list.Slice())
	fmt.Println(list.Slice(1))
	fmt.Println(list.Slice(0, 1))

	fmt.Println(list.MapCan(func(val interface{}) *List {
		fmt.Println(val)
		return New(val, "/")
	}))

	fmt.Println(list.When(func (val interface{}) bool {
		_, ok := val.(int)
		return ok
	}))

	fmt.Println(list.MapCar(func (val interface{}) interface{} {
		if v, ok := val.(int); ok {
			return v * 1000
		} else { return val }
	}))

	list.MapInto(func (val interface{}) interface{} {
		return nil
	})
	fmt.Println(list)

	list.MapRef(func (val *interface{}) {
		*val = 1
	})
	fmt.Println(list)

	fmt.Println("")

	list = New(1)
	fmt.Println(list)

	list = New(1, 2, 3)
	fmt.Println(list)

	list = New([]int{1, 2, 3 ,4 ,5})
	fmt.Println(list)

	list = New([]int{1,2}, []int{3,4})
	fmt.Println(list)
	
	list = New(New(1, 2, 3, 4, 5, 6))
	fmt.Println(list)
	fmt.Println(list.Slice(123, -559))
}
