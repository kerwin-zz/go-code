package main

import "slices"

/*
	给定一个列表 accounts，每个元素 accounts[i] 是一个字符串列表，其中第一个元素 accounts[i][0] 是 名称 (name)，其余元素是 emails 表示该账户的邮箱地址。

	现在，我们想合并这些账户。如果两个账户都有一些共同的邮箱地址，则两个账户必定属于同一个人。请注意，即使两个账户具有相同的名称，它们也可能属于不同的人，因为人们可能具有相同的名称。一个人最初可以拥有任意数量的账户，但其所有账户都具有相同的名称。

	合并账户后，按以下格式返回账户：每个账户的第一个元素是名称，其余元素是 按字符 ASCII 顺序排列 的邮箱地址。账户本身可以以 任意顺序 返回。
*/

/*
	算法：
	1. 把 accounts 中的信息提取到哈希表 emailToIdx 中，key 为邮箱地址，value 为这个邮箱对应的账户下标列表。
	2. 初始化一个长为 n 的全为 false 的布尔数组 vis，用来标记访问过的账户下标。
	3. 遍历 vis，如果 i 没有访问过，即 vis[i] = false，这从 i 开始 dfs。
		a. dfs 之前，创建一个哈希集合 emails，用来保存 dfs 中访问到的邮箱地址。
		b. 开始 dfs。首先标记 vis[i] = true。
		c. 遍历 accounts[i] 的邮箱 email。
		d. 如果 email 在哈希集合 emails 中，则跳过；否则把 email 加入哈希集合 emails。
		e. 遍历 emailToIdx[email]，也就是所有包含改邮箱地址的账户下标 j，如果 j 没有被访问过，即 vis[j] = false，则继续 dfs j。
	4. dfs 结束后，把 emails 中的元素按照字典序从小到大排序，然后和 accounts[i][0] 一起加入答案。
	5. 返回答案。

*/

func accountsMerge(accounts [][]string) (ans [][]string) {
	emailToIdx := make(map[string][]int)
	for i, account := range accounts {
		for _, email := range account[1:] {
			emailToIdx[email] = append(emailToIdx[email], i)
		}
	}

	vis := make([]bool, len(accounts))
	emailSet := make(map[string]struct{}) // 用于收集 dfs 中访问到的邮箱地址
	var dfs func(int)
	dfs = func(i int) {
		vis[i] = true
		for _, email := range accounts[i][1:] {
			if _, ok := emailSet[email]; ok {
				continue
			}
			emailSet[email] = struct{}{}
			for _, j := range emailToIdx[email] { // 遍历所有包含该邮箱地址的账户下标 j
				if !vis[j] { // j 没有被访问过
					dfs(j)
				}
			}
		}
	}

	for i, b := range vis {
		if b {
			continue
		}
		clear(emailSet)
		dfs(i)

		res := make([]string, 1, len(emailSet)+1)
		res[0] = accounts[i][0]
		for email := range emailSet {
			res = append(res, email)
		}
		slices.Sort(res[1:])

		ans = append(ans, res)
	}
	return
}
