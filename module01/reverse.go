package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//	Reverse("cat") => "tac"
//	Reverse("alphabet") => "tebahpla"
func Reverse(word string) string {
	if len(word) == 0 {
		return ""
	}
	rev := ""
	for i := len(word) - 1; i >= 0; i-- {
		rev += string(word[i])
	}
	return rev
}
