package utils

import "fmt"

//Check thorws panic(e) if e not nil
func Check(e error) {
	if !Nil(e) {
		panic(e.Error())
	}
}

//Nil returns true if given interface is nil
func Nil(t interface{}) bool {
	return t == nil
}

func Literal(str string) string {
	return fmt.Sprintf("%q", str)
}
