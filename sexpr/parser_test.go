package sexpr

import (
	"testing"
	"fmt"
)

func TestSexpr(t *testing.T) {
	str := ` 
	
( one 
  two 
  "three"
  " \"four\" "
  (\five 
    \"six 
    (7)
    one(two)three
    (4)(5)
    (8 9)) () (10) 11)
`
	


//KeepQuotes= true
	fmt.Println(str)

	list := Parse(str)
	full, _ := ParseFull(str)
	full2 := DeepClone(full)
	Path(full, 4, 2, "CHANGEEE")
	fmt.Println(full)
	fmt.Println(full2)
	
	fmt.Println(String(list))
	fmt.Println(ToList(list))
	fmt.Println(ParseFull(str))
}
