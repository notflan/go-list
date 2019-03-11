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

func Flat(vl ...[]interface{})  *List {
	ret := New()
	for _, v:=range vl {
		ret.Add(v)
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

func (this *List) Map(f func (interface{})) {
	sl := *this
	for _, v := range sl {
		f(v)
	}
}
func (this *List) MapCar(f func (interface{}) interface{} ) *List {
	sl := *this
	ret := New()
	for _, v := range sl {
		ret.Add(f(v))
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

func (this *List) MapInto(f func (interface{}) interface{}) {
	for i, v := range *this {
		(*this)[i] = f(v)
	}
}

func (this *List) MapRef(f func(*interface{})) {
	sl := *this
	for i,_ := range sl {
		f(&sl[i])
	}
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

func (this *List) MapCan(f func(interface{}) *List) *List {
	over := *this
	tmp := make([]*List, len(over))

	for i, v :=range over {
		tmp[i] = f(v)
	}

	return Nconc(tmp...)
}

func Single(val interface{}) *List {
	return New().Add(val)
}

func (this *List) When(f func(interface{}) bool) *List {
	return this.MapCan(func(car interface{}) *List {
		if f(car) {
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
