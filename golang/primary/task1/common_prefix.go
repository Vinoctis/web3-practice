package main 

import "fmt"

func CommonPrefix(s []string) string {
	if len(s) == 0 {
		return ""
	}
	// 遍历首字符串的字符
	for i:=0;i<len(s[0]);i++ {
		//遍历其他字符串的字符
		for j:=1;j<len(s);j++ {
			//如果长度超出字符串长度，或出现不相同的情况直接返回首字符串当前遍历索引，就是公共前缀
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