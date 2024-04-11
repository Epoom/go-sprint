package sprint

func StrToInt(s string) int {
	x := 1
	nb := 0
	i := 0

	if s[i] == ' ' {
		return 0
	}
	for s[i] == '-' || s[i] == '+' {
		if s[i] == '-' {
			x *= -1
		}
		i++
	}
	for s[i] >= '0' && s[i] <= '9' {
		nb = (nb * 10) + (int(s[i] - '0'))
		i++
	}
	nb *= x
	return nb
}