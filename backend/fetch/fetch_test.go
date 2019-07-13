package fetch

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	result, err := Fetchurl("1")
	if err != nil {
		fmt.Println(err)
	}
	str := string(result)
	str = str[10 : len(str)-2]

	fmt.Println(str)
}