package utils

//获取“=”号的索引
func Substring(str string) int {
	var index int
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "=" {
			index = i
			break
		}
	}
	return index
}

func GetCityid(str string) string {
	rs := []rune(str)
	start := Substring(str) + 1
	end := len(str)
	length := len(rs)

	if start > length {
		return ""
	}

	if end < 0 || end > length {
		return ""
	}
	return string(rs[start:end])
}
