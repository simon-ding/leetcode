package leetcode

func isHappy(n int) bool {
	var seen = make(map[int]bool, 0)
	return happy(n, seen)
}

func happy(n int, seen map[int]bool) bool {
	seen[n] = true
	if n == 1 {
		return true
	}

	n = sum1(n)
	exist := seen[n]
	if exist {
		return false
	}

	return happy(n, seen)
}

func sum1(n int) int {
	sum := 0
	for n > 0 {
		t := n % 10
		n = n / 10
		sum += t * t
	}
	return sum
}
