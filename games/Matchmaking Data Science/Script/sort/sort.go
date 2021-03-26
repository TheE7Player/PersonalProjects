package sort

import "sort"

func SortByFloat(table map[string]float32) map[string]float32 {
	len := len(table)
	temp := make([]float64, 0, len)
	returnVal := make(map[string]float32, len)
	returnTemp := make(map[float32]string, len)

	for k, v := range table {
		temp = append(temp, float64(v))
		returnTemp[v] = k
	}

	sort.Float64s(temp)

	for i := 0; i < len; i++ {
		returnVal[returnTemp[float32(temp[i])]] = float32(temp[i])
	}

	return returnVal
}
