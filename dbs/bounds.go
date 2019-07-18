package dbs

func signum(x float64) int8 {
	if x == 0 {
		return 0
	}
	if x < 0 {
		return -1
	}
	return +1
}
