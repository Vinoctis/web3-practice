package main 

import "fmt"

func CommonPrefix(s []string) string {
	if len(s) == 0 {
		return ""
	}

	for i:=0;i<len(s[0]);i++ {
		for j:=1;j<len(s);j++ {
			if i >= len(s[j]) - 1  || s[j][i] != s[0][i] {
				return s[0][:i]
			}
		}
	}
	return s[0]
}

func main() {
	s := []string{"flower","flow",""}
	fmt.Println(CommonPrefix(s))
}