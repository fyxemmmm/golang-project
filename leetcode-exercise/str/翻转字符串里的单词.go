package str

import "strings"

func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	arr2 := []string{}
	for i := len(arr)-1; i >= 0; i -- {
		if arr[i] != "" {
			arr2 = append(arr2, arr[i])
		}
	}
	return strings.Join(arr2, " ")
}