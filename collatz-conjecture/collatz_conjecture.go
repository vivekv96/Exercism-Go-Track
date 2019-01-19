package collatzconjecture

import "errors"

func CollatzConjecture(input int) (int, error) {
	if input <= 0 {
		return 0, errors.New("error")
	}
	steps := 0
	for {
		if input == 1 {
			break
		}
		if input%2 == 0 {
			input /= 2
			steps++
		} else {
			input = 3*input + 1
			steps++
		}
	}

	return steps, nil
}
