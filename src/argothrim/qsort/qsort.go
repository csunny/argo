package qsort

func QuickSort(values []int) []int {
	if len(values) <= 1{
		return values
	}

	mid := values[0]   // 基准值
	head, tail := 0, len(values) - 1

	// 遍历整个切片
	for i :=1; i <= tail; {
		if values[i] > mid{
			// 小的数据放在基准值左侧
			values[i], values[tail] = values[tail], values[i]
			tail --
		}else {
			values[i], values[head] = values[head], values[i]
			head ++
			i ++
		}
	}
	QuickSort(values[:head])
	QuickSort(values[head+1:])

	return values
}
