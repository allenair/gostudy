package mytool

func Leecode01(numarr []int, target int) []int {
	resArr := make([]int, 0)
	arrMap := make(map[int]int)

	for index, value := range numarr {
		arrMap[value] = index
		if ntindex, ok := arrMap[target-value]; ok {
			resArr = append(resArr, ntindex, index)
		}
	}

	return resArr
}
