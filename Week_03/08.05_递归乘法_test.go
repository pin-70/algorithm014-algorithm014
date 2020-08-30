package Week_03

// 不让用*号。可以转换成求加法。需要注意的是，为了防止递归栈太深，求AB的最小值，用最小值去-1，减少调用栈的深度
func multiply(A int, B int) int {
	if A == 0 || B == 0 {
		return 0
	}
	if A < B {
		return multiply(A-1, B) + B
	}
	return multiply(B-1, A) + A
}

/* 位移方法看不懂
class Solution:
    def multiply(self, A: int, B: int) -> int:
        if B == 1: return A
        if B == 0: return 0
        if B & 1:
            return self.multiply(A<<1, B>>1) + A
        else:
            return self.multiply(A<<1, B>>1)
*/
