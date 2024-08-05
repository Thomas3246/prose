package prose

import "strings"

func JoinWinthCommas(phrases []string) string {
	result := strings.Join(phrases[:len(phrases)-1], ", ")
	result += ", and "
	result += phrases[len(phrases)-1]
	return result
}
