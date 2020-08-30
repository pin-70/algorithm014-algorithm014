package Week_03

func combine(n int, k int) [][]int {
	if n < k || k == 0 {
		return nil
	}
	res := make([][]int, 0)
	for i := 1; i <= n-k+1; i++ {

		res = append(res, genNum(i, k))
	}
	return res
}

func genNum(start int, k int) (res []int) {
	for i := 1; i <= k; i++ {
		res = append(res, start+i-1)
	}
	return
}
