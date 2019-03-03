package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func main() {

	strIntervals := []strInterval{}

	//s := "aaaabbcc"
	s := "abcxyz123"
	//dict := []string{"aaa", "aab", "bc"}
	dict := []string{"abc", "123"}

	addBoldTag(s, dict)
	fmt.Println(merge(strIntervals))

}

func addBoldTag(s string, dict []string) string {

	strIntervals := []strInterval{}
	for _, str := range dict {

		if strings.Contains(s, str) {
			index := -1

			index = strings.Index(s, str)
			strInterval := strInterval{Start: index, End: index + len(str)}
			strIntervals = append(strIntervals, strInterval)
		}

	}

	resultstrIntervals := merge(strIntervals)

	var buffer bytes.Buffer
	strArr := []rune(s)
	i := 0
	for j := 0; j < len(resultstrIntervals); j++ {

		startIndex := resultstrIntervals[j].Start
		endIndex := resultstrIntervals[j].End

		for i <= endIndex && i <= len(s) {

			if i == startIndex {
				buffer.WriteRune('<')
				buffer.WriteRune('b')
				buffer.WriteRune('>')

			} else if i == endIndex {

				buffer.WriteRune('<')
				buffer.WriteRune('/')
				buffer.WriteRune('b')
				buffer.WriteRune('>')

			}
			if i < len(strArr) {
				buffer.WriteRune(strArr[i])
			}
			i++

		}

	}

	fmt.Print(buffer.String())
	return buffer.String()

}

type strInterval struct {
	Start int
	End   int
}

func merge(strIntervals []strInterval) []strInterval {
	if len(strIntervals) == 0 {
		return []strInterval{}
	}

	sort.Slice(strIntervals, func(i, j int) bool { return strIntervals[i].Start < strIntervals[j].Start })
	ret := []strInterval{strIntervals[0]}
	for i := 1; i < len(strIntervals); i++ {
		tail := ret[len(ret)-1]
		if tail.End < strIntervals[i].Start {
			ret = append(ret, strIntervals[i])
		} else if strIntervals[i].End > tail.End {
			ret[len(ret)-1].End = strIntervals[i].End
		}
	}
	return ret
}
