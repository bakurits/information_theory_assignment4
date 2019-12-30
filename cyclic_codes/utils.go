package cyclic_codes

// binPow evaluates (base ^ deg) % rem
func binPow(base int, deg int, rem int) int {
	var res = 1
	for deg > 0 {
		if (deg & 1) > 0 {
			res = int(int64(res) * int64(base) % int64(rem))
		}
		base = int((int64(base) * int64(base)) % int64(rem))
		deg >>= 1
	}
	return res
}

func findInverse(a int, p int) int {
	return binPow(a, p-2, p)
}
