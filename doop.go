package sprint

func Doop(a int, op string, b int) int {
	// start a switch condition that swicthes the value of op
		switch op {
			// if the value of op is "+", switches its value to the value of a added the value of b and returns it
			case "+":
			return a + b
			// if the value of op is "-", switches its value to the value of a subtracked the value of b and returns it
			case "-":
			return a - b
			// if the value of op is "/", switches its value to the value of a divided the value of b and returns it
			case "/":
				// sets an if condition that checks if the value of b equals 0
				if b == 0{
					// when the if condition is met, returns 0
					return 0
				}
			return a / b
			// if the value of op is "*", switches its value to the value of a multiplied the value of b and returns it
			case "*":
			return a * b
			// if the value of op is "%", switches its value to the value of a modulo the value of b and returns it
			case "%":
				// sets an if condition that checks if the value of b equals 0
				if b == 0 {
					// when the if condition is met, returns 0
					return 0
				}
			return a % b
		}
		// if no switch case is met, returns 0
		return 0
	}