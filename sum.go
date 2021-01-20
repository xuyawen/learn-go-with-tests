package integers

func Sum(arr []int) int {
	var sum int
	for _, v := range arr {
		sum += v
	}
	return sum
}

func SumAll(allArr ...[]int) int {
	sum := 0
	for _, arr := range allArr {
		for _, num := range arr {
			sum += num
		}
	}
	return sum
}

func SumAllItem(numbersToSum ...[]int) (sums []int) {
	// lengthOfNumbers := len(numbersToSum)
	// sums = make([]int, lengthOfNumbers)
	// for i, numbers := range numbersToSum {
	// 	sums[i] = Sum(numbers)
	// }
	for _, v := range numbersToSum {
		sums = append(sums, Sum(v))
	}
	return
}

func SumAllTails(numberTailsSum ...[]int) (sums []int) {
	for _, v := range numberTailsSum {
		if len(v) == 0 {
			sums = append(sums, 0)
		} else {
			tail := v[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}
