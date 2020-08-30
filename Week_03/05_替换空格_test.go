package Week_03

import "strings"

// 使用库方法
func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

// 依次判断
func replaceSpace2(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) != " " {
			res = res + string(s[i])
		} else {
			res = res + "%20"
		}
	}
	return res
}
