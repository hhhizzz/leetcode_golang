package _344

//golang真是有毒
func reverseString(s []byte) {
	for i := 0; i < len(s)>>1; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}
