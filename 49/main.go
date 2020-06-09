package _49

import "strconv"

func toCount(str string, count *[26]int) string {
	for i := 0; i < 26; i++ {
		count[i] = 0
	}
	for _, c := range str {
		(*count)[c-'a'] += 1
	}
	result := ""
	for i := range *count {
		if count[i] != 0 {
			result += strconv.Itoa(count[i])
			result += string(i + 'a')
		}
	}
	return result
}

func groupAnagrams(strs []string) [][]string {
	m := map[string]int{}
	count := [26]int{}
	result := make([][]string, 0)
	for _, str := range strs {
		key := toCount(str, &count)
		if _, ok := m[key]; ok {
			result[m[key]] = append(result[m[key]], str)
		} else {
			m[key] = len(result)
			result = append(result, []string{str})
		}
	}
	return result
}
