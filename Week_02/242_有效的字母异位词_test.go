package Week_02

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mp := make(map[uint8]int, 0)
	for i := 0; i < len(s); i++ {
		mp[s[i]]++
		mp[t[i]]--
	}
	for _, v := range mp {
		if v != 0 {
			return false
		}
	}
	return true
}
