package Week_02

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	mp := make(map[[26]int][]string)

	for _, s := range strs {
		letters := [26]int{}
		for _, c := range s {
			letters[c-'a']++
		}
		mp[letters] = append(mp[letters], s)
	}

	for _, v := range mp {
		res = append(res, v)
	}
	return res
}
