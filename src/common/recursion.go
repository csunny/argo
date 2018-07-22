package common


// 斐波拉切数列
func Fab(n int) int {
	if n < 2 {
		return n
	}

	return Fab(n-1) + Fab(n-2)
}


// 引申 hash表实现斐波拉切数列

func Fab_Hash(n int) int  {
	result := make(map[int]int)

	result[0] = 0
	result[1] = 1

	for i:=2; i<=n; i++{
		result[i] = result[i-1] + result[i-2]
	}

	return result[n]
}
