package sprint

func FindDivided(from, to, divisor int) int{
for i := from ; i < to ; i++ {
	if i % divisor == 0 {
		return i
	}
}	
	return -1
}