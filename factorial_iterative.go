package sprint


func FactorialIterative(n int) int {

f := 1

if n < 0 {
	return 0
}
for i := 0; i <= n; i++ {
	f = f * i
}
return f
}