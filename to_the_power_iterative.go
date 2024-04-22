package sprint

func ToThePowerIterative(n int, power int) int {

res := 1

if power < 0 {
	return 0
}
if power == 0 {
	return 1
}
for power > 0 {
	res *= n
	power--
}
return res
}