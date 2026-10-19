package main

func rotateRunes(s string, shift int) string {
	r := []rune(s)
	n := len(r)
	if n == 0 {
		return ""
	}

	k := shift % n
	if k < 0 {
		k += n
	}

	out := make([]rune, 0, n)
	out = append(out, r[k:]...)
	out = append(out, r[:k]...)
	return string(out)
}
