package main

import "fmt"

func CombN(n int) []string {
	var str []string

	switch n {
	case 1:
		for a := 0; a <= 9; a++ {
			str = append(str, fmt.Sprintf("%d", a))
		}
	case 2:
		for a := 0; a <= 8; a++ {
			for b := a + 1; b <= 9; b++ {
				combination := a*10 + b
				str = append(str, fmt.Sprintf("%02d", combination))
			}
		}
	case 3:
		for a := 0; a <= 7; a++ {
			for b := a + 1; b <= 8; b++ {
				for c := b + 1; c <= 9; c++ {
					combination := a*100 + b*10 + c
					str = append(str, fmt.Sprintf("%03d", combination))
				}
			}
		}
	case 4:
		for a := 0; a <= 6; a++ {
			for b := a + 1; b <= 7; b++ {
				for c := b + 1; c <= 8; c++ {
					for d := c+1; d <= 9; d++ {
						combination := a*1000 + b*100 + c*10 + d
						str = append(str, fmt.Sprintf("%04d", combination))
					}
				}
			}
		}
	case 5:
		for a := 0; a <= 5; a++ {
			for b := a + 1; b <= 6; b++ {
				for c := b + 1; c <= 7; c++ {
					for d := c+1; d <= 8; d++ {
						for e := d+1; e <= 9; e++ {
							combination := a*10000 + b*1000 + c*100 + d*10 + e
							str = append(str, fmt.Sprintf("%05d", combination))
						}
					}
				}
			}
		}
	case 6:
		for a := 0; a <= 4; a++ {
			for b := a + 1; b <= 5; b++ {
				for c := b + 1; c <= 6; c++ {
					for d := c+1; d <= 7; d++ {
						for e := d+1; e <= 8; e++ {
							for f := e+1; f <= 9; f++ {
								combination := a*100000 + b*10000 + c*1000 + d*100 + e*10 + f
								str = append(str, fmt.Sprintf("%06d", combination))
							}
						}
					}
				}
			}
		}
	case 7:
		for a := 0; a <= 3; a++ {
			for b := a + 1; b <= 4; b++ {
				for c := b + 1; c <= 5; c++ {
					for d := c+1; d <= 6; d++ {
						for e := d+1; e <= 7; e++ {
							for f := e+1; f <= 8; f++ {
								for g := f+1; g <= 9; g++ {
									combination := a*1000000 + b*100000 + c*10000 + d*1000 + e*100 + f*10 + g
									str = append(str, fmt.Sprintf("%07d", combination))
								}
							}
						}
					}
				}
			}
		}
	case 8:
		for a := 0; a <= 2; a++ {
			for b := a + 1; b <= 3; b++ {
				for c := b + 1; c <= 4; c++ {
					for d := c+1; d <= 5; d++ {
						for e := d+1; e <= 6; e++ {
							for f := e+1; f <= 7; f++ {
								for g := f+1; g <= 8; g++ {
									for h := g+1; h <= 9; h++ {
										combination := a*10000000 + b*1000000 + c*100000 + d*10000 + e*1000 + f*100 + g*10 + h
										str = append(str, fmt.Sprintf("%08d", combination))
									}
								}
							}
						}
					}
				}
			}
		}
	case 9:
		for a := 0; a <= 1; a++ {
			for b := a + 1; b <= 2; b++ {
				for c := b + 1; c <= 3; c++ {
					for d := c+1; d <= 4; d++ {
						for e := d+1; e <= 5; e++ {
							for f := e+1; f <= 6; f++ {
								for g := f+1; g <= 7; g++ {
									for h := g+1; h <= 8; h++ {
										for i := h+1; i <= 9; i++ {
											combination := a*100000000 + b*10000000 + c*1000000 + d*100000 + e*10000 + f*1000 + g*100 + h*10 + i
											str = append(str, fmt.Sprintf("%09d", combination))
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return str
}

func main() {
	combinations := CombN(9)

	fmt.Println("Array of Combinations:")
	fmt.Println(combinations)
}