package scrapbox

import "fmt"

type SBNode interface {
	fmt.Stringer
	IsValid(text string) bool
}
