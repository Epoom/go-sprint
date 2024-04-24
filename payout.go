package sprint

func Payout(amount int, denominations []int) (payout []int) {
for i := 0; i < len(denominations); i++ {
	for j := i + 1; j < len(denominations); j ++ {
		if denominations[i] < denominations[j] {
			denominations[i], denominations[j] = denominations[j], denominations[i]
		}
	}
}

res := make ([]int, 0)
for _, denomination := range denominations {
	for amount >= denomination {
		res = append(res, denomination)
		amount -= denomination
	}
}
if amount == 0 {
	return res
}
return []int{}
}
