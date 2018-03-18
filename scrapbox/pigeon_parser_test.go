package scrapbox

import (
	"fmt"
	"testing"
)

func TestPigeonParser(t *testing.T) {
	ret, err := Parse("test", []byte("hogehoge"))
	if err != nil {
		panic(err)
	}

	rets, ok := ret.([]interface{})
	if !ok {
		fmt.Println("ret is not []interface{}")
	}

	for _, r := range rets {
		if bytesR, ok := r.([]byte); ok {
			fmt.Println(string(bytesR))
		}
	}

	//switch r := ret.(type) {
	//case string:
	//	fmt.Println("ret is string")
	//	fmt.Println(r)
	//case byte:
	//	fmt.Println("ret is byte")
	//	fmt.Println(r)
	//case []byte:
	//	fmt.Println("ret is []byte")
	//	fmt.Println(r)
	//case []rune:
	//	fmt.Println("ret is []rune")
	//	fmt.Println(r)
	//default:
	//	fmt.Println(reflect.TypeOf(r))
	//	fmt.Println(r)
	//}
}
