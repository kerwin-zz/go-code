package main

/*
	操作1： 00->10	操作2：10->01
	提示 1：
		答案不会包含 00（连续的 0）。
		证明：反证法。假设答案包含 00，我们可以通过操作 1 将其变为 10，从而得到更大的答案。所以
			 答案不会包含 00。
	提示 2：
		答案最多包含一个 0。
		证明：反证法。假设至少有两个 0，我们随意选择其中两个 0，由提示 1 可知这两个 0 不相邻。例
			 如 10110，我们可以通过操作 2 将右边的 0 移动到第一个 0 的右边，即 10011，然后通
			 过操作 1 将其变成 11011。由于昨天更高位的 0 变成了1，所以我们的到了比 10110 更大
			 的答案。一般地，在有多个 0 的情况下，总是可以通过操作 2 让最高位的 0 的右侧也是0，
			 然后通过操作 1 让最高位的 0 变成 1，从而得到更大的答案。所以答案最多包含一个 0。
	提示 3：
		如果 binary 全是 1，直接返回即可。

		如果 binary 中有 0，由于操作 1 和操作 2 的结果都包含 0，所以我们无法把所有的 0 都变成 1。
		结合提示 2，最终答案会恰好包含一个 0。

		此外，提示 2 相当于给出了一个让二进制更大的方案：只要还有两个 0，那么用操作 2 把右边的 0 往
		左移，当出现 00 时就通过操作 1 把左边的 0 变成 1，这会让二进制更大。

		设 binary 从左到右第一个 0 的下标为 i，为了得到更大的二进制，下标在 [i, n-1] 中的 1会随着
		0 的左移被挤到 binary 的末尾。例如
					101010 -> 100110 -> 110110 -> 110011 -> 111011
                         操作 2      操作 1    操作 2     操作 1
		或者
					101010 -> 100011 -> 111011
						 操作 2      操作 1

		注意 101010 -> 111110 是无法做到的，在末尾 0 不移动的情况下，我们无法把前面的 0 全部变成 1。
		一般地，在有多个不相邻 0 的情况下，不移动末尾 0 又能把前面所有 0 变成 1 是做不到的，因为 0 只能
		左移不能右移。

		一般地，设 [i, n-1] 中有 cnt1 个 1，那么答案中唯一的 0 的下标为 n-1-cnt1。
*/

import "strings"

func maximumBinaryString(binary string) string {
	i := strings.Index(binary, "0")
	if i < 0 { // 全是 '1'
		return binary
	}
	cnt1 := strings.Count(binary[i:], "1") // 统计 binary[i:] 中 '1' 的个数
	return strings.Repeat("1", len(binary)-1-cnt1) + "0" + strings.Repeat("1", cnt1)
}