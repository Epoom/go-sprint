package sprint

func StrToInt(s string) int {
	x := 1
	nb := 0
	i := 0

	if len(s) == 0 || s[i] == ' ' {
		return 0
	}
	for s[i] == '-' || s[i] == '+' {
		if s[i] == '-' {
			x *= -1
		}
		i++
	}
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		nb = (nb * 10) + (int(s[i] - '0'))
		i++
	}
	if i < len(s) {
		return 0
	}
	nb *= x
	return nb
}