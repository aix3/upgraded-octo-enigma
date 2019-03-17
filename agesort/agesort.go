package agesort

func agesort(ages []int, max int) {
	timesOfAge := make([]int, max)
	for _, v := range ages {
		timesOfAge[v]++
	}

	var index int
	for i := 0; i < max; i++ {
		times := timesOfAge[i]
		for ; times > 0; times-- {
			ages[index] = i
			index++
		}
	}
}
