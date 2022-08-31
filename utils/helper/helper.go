package helper

import "regexp"

func Filter(s []string, url string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		matched, _ := regexp.MatchString(url, str)
		if _, ok := inResult[str]; !ok && !matched {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}
