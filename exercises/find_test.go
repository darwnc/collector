package exercises

import (
	"fmt"
	"testing"
)

func TestFindString(t *testing.T) {
	// "barfoofoobarthefoobarman"
	// ["bar","foo","the"]
	result := findSubstring("barfoofoo我barthe我foobarman", []string{"bar", "foo我", "the我"})
	// fullString := strings.Join([]string{"1", "12"}, "")
	fmt.Println("result=", result)
}
