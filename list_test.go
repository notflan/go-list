package list

import (
	"testing"
	"fmt"
)

func TestAssoc(t *testing.T) {

	list := NewAlist("key",1,"key2",2,3,4.4,5,nil,6,nil,"key3",New(1,2,3),"key2",nil,"key4", []interface{} {"hello", "world"}, "K")

	fmt.Println(list.Assoc("key3"))
	fmt.Println(list.Assoc("key2"))
	fmt.Println(list.Assoc("K"))
	fmt.Println(list.Assoc("k"))
	fmt.Println(list.Assoc("five", func(a interface{}, b interface{}) bool {
			return a != b
	}, func(a interface{}) interface{} {
			if a == "key" { return "five"
			} else { return a }
	}))
}

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

	fmt.Println(list.MapCar(func (i int, val interface{}) interface{} {
		if v, ok := val.(int); ok {
			return v * i
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

	list = New([]int{-1, 0, 1, 2, 3 ,4 ,5})
	list.RemoveInto(0,1,-1)
	list.Remove(0,1)
	fmt.Println("->", list.ToStrings(func (_ *List,_  int, v interface{}) string { return fmt.Sprintf("=%v", v) }))
	fmt.Println(list)
	list.Map(func(index int, value interface{}) {
		fmt.Println(index, value)
	})
	list.Map(func(value interface{}) {
		fmt.Println("?", value)
	})

	list = New([]int{1,2}, []int{3,4})
	fmt.Println(list)
	
	list = New(New(1, 2, 3, 4, 5, 6))
	fmt.Println(list)
	fmt.Println(list.Slice(123, -559))

	list = Flat([]int{1,2}, []int{3,4,5}, 6, 7, "eight")
	fmt.Println(list)
}
