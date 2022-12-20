package compare

func Compare[N Ordered](a, b N) int {
	switch {
	case a > b:
		return +1
	case a < b:
		return -1
	default:
		return 0
	}
}
