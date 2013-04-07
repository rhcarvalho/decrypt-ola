package caesar

func Shift(n int) func(rune) rune {
	return func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+rune(n))%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+rune(n))%26
		}
		return r
	}
}

var ROT13 = Shift(13)

