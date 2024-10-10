package sum

func SumArr(numbers [5]int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}

func SumShort(numbers []int) int {
	sum := 0
	if len(numbers) > 0 {
		for _, num := range numbers {
			sum += num
		}
	}
	return sum
}

func SumAllSlices(numsToSum ...[]int) []int {
	sliceSum := 0
	var allSums []int
	for _, slice := range numsToSum {
		if len(slice) > 0 {
			sliceSum = SumShort(slice)
			allSums = append(allSums, sliceSum)
		}
	}
	return allSums
}

// Input: 1 or more slice | Output: a slice with the sum of elemens of each input slice
func SumAllSlicesAlt(numbersToSum ...[]int) []int {
	var inputsQty int = len(numbersToSum)
	sumsArr := make([]int, inputsQty)
	for idx, nums := range numbersToSum {
		if len(nums) > 0 {
			sumsArr[idx] = SumShort(nums)
		}
	}
	return sumsArr
}

// Input: One or more slice | Output: a slice, each element is the sum of the given slice/s input except the first element of the array
func SumAllTails(slices ...[]int) []int {
	var sums []int
	for i := 0; i < len(slices); i++ {
		if len(slices[i]) == 0 {
			sums = append(sums, 0)
		} else {
			tail := slices[i][1:]
			if len(tail) == 0 {
				sums = append(sums, 0)
			} else {
				sums = append(sums, SumShort(tail))
			}
		}
	}
	return sums
}
