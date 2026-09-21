func isPalindrome(s string) bool {
	ss := strings.ToLower(s)
	i, j := 0, len(ss) - 1

	for i <= j {
		if (ss[i] < 'a' || ss[i] > 'z') && (ss[i] < '0' || ss[i] > '9') {
			i++
		} else if (ss[j] < 'a' || ss[j] > 'z') && (ss[j] < '0' || ss[j] > '9') {
			j--
		} else if ss[i] == ss[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}
