package exercises

import (
	"fmt"
	"testing"
)

func TestFindString(t *testing.T) {
	result := findSubstring("barfoothefoobarman", []string{"foo", "bar"})
	// fullString := strings.Join([]string{"1", "12"}, "")
	fmt.Println("result=", result)
}
