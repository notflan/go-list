package list

import (
	"fmt"
	"reflect"
)

type List  []interface{}

func New(initial ...interface{}) *List {
	if initial == nil || len(initial)<1 {
		val := make(List, 0)
		return &val
	} else if len(initial) == 1 {
		switch initial[0].(type) {
			case []interface{}:
				return New().Add(initial[0].([]interface{})...)
			case *List:
				return (initial[0].(*List)).Clone()
			default:
				if reflect.TypeOf(initial[0]).Kind() == reflect.Slice {
					vr := reflect.ValueOf(initial[0])
					val := make([]interface{}, vr.Len())
					for i :=0;i<vr.Len();i++ {
						val[i] = vr.Index(i).Interface()
					}

					return New(val)
				} else {
					return New().Add(initial[0])
				}
		}
	} else {
		return New().Add(initial...)
	}
}

func Flat(vl ...interface{})  *List {
	ret := New()
	for _, v:=range vl {
		switch v.(type) {
			case []interface{}:
				ret.Add(v.([]interface{})...)
			default:
				if reflect.TypeOf(v).Kind() == reflect.Slice {
					value := reflect.ValueOf(v)
					add := make([]interface{}, value.Len())
					for i:=0;i<value.Len();i++ {
						add[i] = value.Index(i).Interface()
					}
					ret.Add(add...)
				} else {
					ret.Add(v)
				}
		}
		//ret.Add(v)
	}
	return ret
}

func lslice(v interface{}) (*List, bool) {
	if v == nil { return nil, false }
	switch v.(type) {
		case *List:
			return v.(*List), true
		default:
			if reflect.TypeOf(v).Kind() == reflect.Slice {
				rt := New()
				val := reflect.ValueOf(v)
				for i:=0;i<val.Len();i++ {
					rt.Add(val.Index(i).Interface())
				}
				return rt, true
			}
	}
	return nil, false
}

func NewAlist(from ...interface{}) *List {
	if len(from) == 0 {
		return New()
	} else if (len(from) % 2) != 0 {
		from = append(from, nil)
	}

	ret := New()
	for i:=0;i<len(from);i+=2 {
		key := from[i]
		value := from[i+1]

		slist, ok := lslice(value)
		if ok {
			ret.Add(New(append([]interface{}{ key }, slist.Slice()...)))
		} else {
			ret.Add(New(key, value))
		}
	}
	return ret
}

func (this *List) Add(vl ...interface{}) *List {
	*this = append(*this, vl...)
	return this
}

func (this *List) Push(vl ...interface{}) *List {
	*this = append(vl, (*this)...)
	return this
}

func (this *List) RangeCheck(index int) bool {
	return index<len(*this) && index >= 0
}

func (this *List) Remove(indecies ...int) *List {
	a := *this

	for _, index := range indecies {
		index = this.Circle(index)
		a = append(a[:index], a[index+1:]...)
	}

	*this = a

	return this
}

func (this *List) RemoveInto(indecies ...int) *List {
	id := make(map[int]bool)
	m := this.Len()
	for _,v := range indecies {
		id[circleSafe(v, m)] = true
	}
	*this = *(this.When(func (index int, _ interface{}) bool {
		_, ok := id[index]
		return !ok
	}))
	return this
}

func (this *List) Map(f interface{}) {
	sl := *this
	g := func(index int, value interface{}) {
		f.(func (int, interface{}))(index,value)
	}
	switch f.(type) {
		case func(interface{}):
			g = func(_ int, value interface{}) {
				f.(func (interface{}))(value)
			}
	}
	for i, v := range sl {
		g(i, v)
	}
}
func (this *List) MapCar(f interface{}) *List {
	sl := *this
	ret := New()
	g := func(index int, value interface{}) interface{} {
		return f.(func(int, interface{}) interface{})(index,value)
	}
	switch f.(type) {
		case func(interface{}) interface{}:
			g = func(_ int, value interface{}) interface{} {
				return f.(func (interface{}) interface{})(value)
			}
	}
	for i, v := range sl {
		ret.Add(g(i, v))
	}
	return ret
}
func (this *List) Slice(vl ...int) []interface{} {
	l := *this

	if vl == nil || len(vl)<1 {
		return l
	} else  if len(vl)<2 {
		
		return l[circleSafe(vl[0], len(l)):]
	} else {
		return l[circleSafe(vl[0], len(l)):circleSafe(vl[1], len(l))]
	}
}

func (this *List) Len() int {
	return len(*this)
}

func (this *List) Cap() int {
	return cap(*this)
}

func circleSafe(index int, max int) int {
	if max <= 0 {
		return 0
	}
	if index < 0 {
		return circleSafe(max+index, max)
	} else {
		return index % max
	}
}

func (this *List) Circle(index int) int {
	return circleSafe(index, this.Len())
}

func (this *List) Get(index int) interface{} {
	index = this.Circle(index)
	if this.RangeCheck(index) {
		return (*this)[index]
	} else {
		return nil
	}
}

func (this *List) Set(index int, val interface{}) {
	index = this.Circle(index)
	if this.RangeCheck(index) {
		(*this)[index] = val
	}
}

func (this *List) MapInto(f interface{}) {
	for i, v := range *this {
		switch f.(type) {
			case func(int,interface{})interface{}:
				(*this)[i] = f.(func(int,interface{}) interface{})(i, v)
			default:
				(*this)[i] = f.(func(interface{}) interface{})(v)
		}
	}
}

func (this *List) MapRef(f interface{}) {
	sl := *this
	for i,_ := range sl {
		switch f.(type) {
			case func(int, *interface{}):
				f.(func(int, *interface{}))(i, &sl[i])
			default:
				f.(func(*interface{}))(&sl[i])
		}
	}
}

func (this *List) Assoc(akey interface{}, eq ...interface{}) (tuple struct { Key, Value interface{} }, ok bool) {
	ok = false
	comp := func(a interface{}, b interface{}) bool {
		return a == b
	}
	key := func(_ int, a interface{}) interface{} {
		return a
	}
	if len(eq)>0 {
		if eq[0] != nil { comp = eq[0].(func(interface{}, interface{}) bool) }
		if len(eq)>1 {
			if eq[1] != nil {
				switch eq[1].(type) {
					case func(int,interface{})interface{}:
						key = eq[1].(func(int,interface{})interface{})
					default:
						key = func(_ int, a interface{}) interface{} {
							return eq[1].(func(interface{})interface{})(a)
						}
				}
			}
		}
	}
	this.Map(func (i int, v interface{}) {
		if ok {
			return
		}
		var slice []interface{}
		switch v.(type) {
			case *List:
				slice = v.(*List).Slice()
			default:
				if reflect.TypeOf(v).Kind() == reflect.Slice {
					value := reflect.ValueOf(v)
					slice = make([]interface{}, value.Len())
					for j :=0;j<value.Len();j++ {
						slice[j] = value.Index(j).Interface()
					}
				} else { 
					slice = []interface{}{v}
				}
		}
		if len(slice)==0 {
			return
		}
		compk := key(i, slice[0])
		switch len(slice) {
			case 0: 
				return
			case 1:
				tuple.Key = slice[0]
				tuple.Value = nil
			case 2:
				tuple.Key = slice[0]
				tuple.Value = slice[1]
			default:
				tuple.Key = slice[0]
				tuple.Value = New(slice[1:])
		}
		if comp(compk, akey) {
			ok = true
		}
	})
	return
}

func (this *List) AddTuple(key interface{}, value interface{}) {
	this.Add(New(key,value))
}

func (this *List) PushTuple(key, value interface{}) {
	this.Push(New(key,value))
}

func Nconc(lists ...*List) *List {
	ret := New()
	for _, v := range lists {
		if v!=nil {
			ret.Add(v.Slice()...)
		}
	}
	return ret
}

func (this *List) MapCan(f interface{}) *List {
	over := *this
	tmp := make([]*List, len(over))

	for i, v :=range over {
		switch f.(type) {
			case func(interface{}) *List:
				tmp[i] = f.(func(interface{}) *List)(v)
			default:
				tmp[i] = f.(func(int,interface{}) *List)(i,v)
		}
	}

	return Nconc(tmp...)
}

func Single(val interface{}) *List {
	return New().Add(val)
}

func (this *List) When(f interface{}) *List {
	g := func(i int, v interface{}) bool {
		return f.(func (int,interface{}) bool)(i,v)
	}
	switch f.(type) {
		case func(interface{}) bool:
			g = func(_ int, v interface{}) bool {
				return f.(func (interface{}) bool)(v)
			}
	}
	return this.MapCan(func(i int, car interface{}) *List {
		
		if g(i, car) {
			return Single(car)
		} else { return nil }
	})
}

func (this *List) Clone() *List {
	return this.MapCar(func(i interface{}) interface{} { return i })
}

func (this *List) Clear() {
	*this = make([]interface{},0)
}

func (this *List) Car() interface{} {
	if this.Len()>0 {
		return (*this)[0]
	} else { return nil }
}

func (this *List) Cdr() *List {
	if this.Len()>1 {
		return New(this.Slice(1)) 
	} else {
		return New()
	}
}

func (this *List) String() string {
	return fmt.Sprintf("%v", *this)
}
