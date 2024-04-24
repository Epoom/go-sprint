package sprint

func FromRoman(s string) int {
romanMap := map[byte]int {
	'I'; 1,
	'V'; 5,
	'X'; 10,
	'L'; 50,
	'C'; 100,
	'D'; 500,
	'M'; 1000,
}

res := 0
prev := 0

for i := len(s) - 1; i >= 0; i-- {
	value := romanMap[s[i]]
	if value < prev {
		res -= value
	} else {
		res += value
	}
	prev = value
}
return res
}