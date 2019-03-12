package sexpr

import (
	"github.com/notflan/go-list"
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

type ValueType int

const (
	VList ValueType = iota
	VString
)

type Value struct {
	t ValueType
	list *list.List
	str string
}

func (a Value) Value() interface{} {
	switch a.t {
		case VList:
			return a.list
		case VString:
			return a.str
	}
	return nil
}

func (a Value) Type() ValueType {
	return a.t
}

func (a Value) SE(se *list.List) Value {
	a.t = VList
	a.list =  se
	return a
}

func (a Value) Str(str string) Value {
	a.t = VString
	a.str = str
	return a
}

func addString(output *list.List, str string){
	output.Add(Value{}.Str(str))
}

func addList(output *list.List, lis *list.List) {
	output.Add(Value{}.SE(lis))
}

func parse(output *list.List, str *[]rune) {
	token := list.New()

	inStr := false
	inEscape := false

	for len(*str)>0 {
		char := (*str)[0]
		*str = (*str)[1:]
		if inEscape {
			inEscape = false
			token.Add(char)
		} else if inStr {
			if char == '"' {
				inStr = false
			} else { token.Add(char) }
		} else if char == ' ' {
			str := string(token.ToRunes())
			token.Clear()
			if len(str)>0 { addString(output, str) }
		} else if char == '(' {
			//Opens a new sexpr
			sub := list.New()
			parse(sub, str)
			addList(output, sub)
		} else if char == ')' {
			//Close this
			break
		} else if char == '"' {
			inStr = true
		} else if char == '\\' {
			inEscape = true
		} else {
			token.Add(char)
		}
	}
	if token.Len() >0 {
		str := string(token.ToRunes())
		token.Clear()
		addString(output, str)
	}
}

func Parse(input string) *list.List {
	valid := regexp.MustCompile(`^\s*\(.*\)\s*$`)

	if !valid.MatchString(input) {
		return nil
	}

	runes := []rune(strings.TrimSpace(input))
	runes = runes[1:len(runes)-1]
	ret := list.New()
	parse(ret, &runes)
	return ret
}


func shouldEscape(str string) bool {
	whiteSpaces := regexp.MustCompile(`\s`)

	return whiteSpaces.MatchString(str)
}

func escape(str string) string {
	return strings.Replace(strings.Replace(str,"\\", "\\\\", -1), "\"", "\\\"", -1)
}

func write(buf *bytes.Buffer, li *list.List) {
	buf.WriteRune('(')
	if li == nil {
		panic("oops")
	}
	for i, car := range li.Slice() {
		value := car.(Value)
		if i != 0 {
			buf.WriteRune(' ')
		}
		if value.t == VString {
			if shouldEscape(value.str) {
				fmt.Fprintf(buf, `"%s"`, escape(value.str))
			} else {
				buf.WriteString(escape(value.str))
			}
		} else {
			write(buf, value.list)
		}
	}
	buf.WriteRune(')')
}

func String(li *list.List) string {
	var buffer bytes.Buffer
	write(&buffer, li)
	return buffer.String()
}

func ToList(ps *list.List) *list.List {
	return ps.MapCan(func (v interface{}) *list.List {
		val := v.(Value)
		switch val.t {
			case VString:
				return list.Single(val.str)
			case VList:
				return list.Single(ToList(val.list))
		}
		return nil
	})
}

func ParseFull(str string) (*list.List, bool) {
	proto := Parse(str)
	if proto==nil {
		return nil, false
	}

	return ToList(proto), true
}
