package list

func (l *List) ToRunes(fails ...func(*List, int, interface{}) rune) []rune {
	ret := make([]rune, l.Len())
	
	l.Map(func (i int, v interface{}) {
		vl, ok := v.(rune)
		if !ok && len(fails)>0 {
			vl = fails[0](l, i, v)
		} else if !ok { panic("Type casting error to rune") }
		ret[i] = vl
	})
	return ret
}

func (l *List) ToInts(fails ...func(*List, int, interface{}) int) []int {
	ret := make([]int, l.Len())
	
	l.Map(func (i int, v interface{}) {
		vl, ok := v.(int)
		if !ok && len(fails)>0 {
			vl = fails[0](l, i, v)
		} else if !ok { panic("Type casting error to int") }
		ret[i] = vl
	})
	return ret
}

func (l *List) ToStrings(fails ...func(*List, int, interface{}) string) []string {
	ret := make([]string, l.Len())
	
	l.Map(func (i int, v interface{}) {
		vl, ok := v.(string)
		if !ok && len(fails)>0 {
			vl = fails[0](l, i, v)
		} else if !ok { panic("Type casting error to string") }
		ret[i] = vl
	})
	return ret
}

func (l *List) ToFloat32s(fails ...func(*List, int, interface{}) float32) []float32 {
	ret := make([]float32, l.Len())
	
	l.Map(func (i int, v interface{}) {
		vl, ok := v.(float32)
		if !ok && len(fails)>0 {
			vl = fails[0](l, i, v)
		} else if !ok { panic("Type casting error to float32") }
		ret[i] = vl
	})
	return ret
}

func (l *List) ToFloat64s(fails ...func(*List, int, interface{}) float64) []float64 {
	ret := make([]float64, l.Len())
	
	l.Map(func (i int, v interface{}) {
		vl, ok := v.(float64)
		if !ok && len(fails)>0 {
			vl = fails[0](l, i, v)
		} else if !ok { panic("Type casting error to float64") }
		ret[i] = vl
	})
	return ret
}

func (l *List) ToBytes(fails ...func(*List, int, interface{}) byte) []byte {
	ret := make([]byte, l.Len())
	
	l.Map(func (i int, v interface{}) {
		vl, ok := v.(byte)
		if !ok && len(fails)>0 {
			vl = fails[0](l, i, v)
		} else if !ok { panic("Type casting error to byte") }
		ret[i] = vl
	})
	return ret
}

func (l *List) ToUints(fails ...func(*List, int, interface{}) uint) []uint {
	ret := make([]uint, l.Len())
	
	l.Map(func (i int, v interface{}) {
		vl, ok := v.(uint)
		if !ok && len(fails)>0 {
			vl = fails[0](l, i, v)
		} else if !ok { panic("Type casting error to uint") }
		ret[i] = vl
	})
	return ret
}

func (l *List) ToInt8s(fails ...func(*List, int, interface{}) int8) []int8 {
	ret := make([]int8, l.Len())
	
	l.Map(func (i int, v interface{}) {
		vl, ok := v.(int8)
		if !ok && len(fails)>0 {
			vl = fails[0](l, i, v)
		} else if !ok { panic("Type casting error to int8") }
		ret[i] = vl
	})
	return ret
}
func (l *List) ToUint8s(fails ...func(*List, int, interface{}) uint8) []uint8 {
	ret := make([]uint8, l.Len())
	
	l.Map(func (i int, v interface{}) {
		vl, ok := v.(uint8)
		if !ok && len(fails)>0 {
			vl = fails[0](l, i, v)
		} else if !ok { panic("Type casting error to uint8") }
		ret[i] = vl
	})
	return ret
}

func (l *List) ToInt16s(fails ...func(*List, int, interface{}) int16) []int16 {
	ret := make([]int16, l.Len())
	
	l.Map(func (i int, v interface{}) {
		vl, ok := v.(int16)
		if !ok && len(fails)>0 {
			vl = fails[0](l, i, v)
		} else if !ok { panic("Type casting error to int16") }
		ret[i] = vl
	})
	return ret
}
func (l *List) ToUint16s(fails ...func(*List, int, interface{}) uint16) []uint16 {
	ret := make([]uint16, l.Len())
	
	l.Map(func (i int, v interface{}) {
		vl, ok := v.(uint16)
		if !ok && len(fails)>0 {
			vl = fails[0](l, i, v)
		} else if !ok { panic("Type casting error to uint16") }
		ret[i] = vl
	})
	return ret
}

func (l *List) ToInt32s(fails ...func(*List, int, interface{}) int32) []int32 {
	ret := make([]int32, l.Len())
	
	l.Map(func (i int, v interface{}) {
		vl, ok := v.(int32)
		if !ok && len(fails)>0 {
			vl = fails[0](l, i, v)
		} else if !ok { panic("Type casting error to int32") }
		ret[i] = vl
	})
	return ret
}
func (l *List) ToUint32s(fails ...func(*List, int, interface{}) uint32) []uint32 {
	ret := make([]uint32, l.Len())
	
	l.Map(func (i int, v interface{}) {
		vl, ok := v.(uint32)
		if !ok && len(fails)>0 {
			vl = fails[0](l, i, v)
		} else if !ok { panic("Type casting error to uint32") }
		ret[i] = vl
	})
	return ret
}

func (l *List) ToInt64s(fails ...func(*List, int, interface{}) int64) []int64 {
	ret := make([]int64, l.Len())
	
	l.Map(func (i int, v interface{}) {
		vl, ok := v.(int64)
		if !ok && len(fails)>0 {
			vl = fails[0](l, i, v)
		} else if !ok { panic("Type casting error to int64") }
		ret[i] = vl
	})
	return ret
}
func (l *List) ToUint64s(fails ...func(*List, int, interface{}) uint64) []uint64 {
	ret := make([]uint64, l.Len())
	
	l.Map(func (i int, v interface{}) {
		vl, ok := v.(uint64)
		if !ok && len(fails)>0 {
			vl = fails[0](l, i, v)
		} else if !ok { panic("Type casting error to uint64") }
		ret[i] = vl
	})
	return ret
}
