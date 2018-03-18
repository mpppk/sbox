package main

import (
	"fmt"

	"github.com/mpppk/sbox/utl"
)

func main() {
	link, err := utl.NewSBLink("[http://yahoo.co.jp Yahoo!]", "https://scrapbox.io/", "niboshi")
	if err != nil {
		panic(err)
	}
	fmt.Println(link)
}
