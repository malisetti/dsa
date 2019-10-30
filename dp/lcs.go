// Package dp is dynamic programming
package dp

// Lcs returns the longest common subsequence from str1 and str2
// TODO: change this to dp
func Lcs(str1, str2 string) string {
	switch {
	case (len(str1) == 0 || len(str2) == 0):
		return ""
	case str1[len(str1)-1] == str2[len(str2)-1]:
		return Lcs(str1[0:len(str1)-1], str2[0:len(str2)-1]) + string(str1[len(str1)-1])
	default:
		l1 := Lcs(str1[0:len(str1)-1], str2)
		l2 := Lcs(str1, str2[0:len(str2)-1])
		if len(l1) > len(l2) {
			return l1
		}
		return l2
	}
}
