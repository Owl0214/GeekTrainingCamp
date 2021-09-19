package day01

import (
	"fmt"
	"testing"
)

func TestStringArrayReplace(t *testing.T) {
	strArr := []string{"I", "am", "stupid", "and", "weak"}
	fmt.Println(strArr)
	replaceMap := make(map[string]string, 2)
	replaceMap["stupid"] = "smart"
	replaceMap["weak"] = "strong"
	replace(strArr, replaceMap)
	fmt.Println(strArr)
}

func replace(src []string, replaceMap map[string]string) {
	for i, val := range src {
		if _, ok := replaceMap[val]; ok {
			src[i] = replaceMap[val]
		}
	}
}

func TestSwitch(t *testing.T) {
	s := 1
	switch s {
	case 1:
		fallthrough
	default:
		fmt.Println("default")
	}
}
