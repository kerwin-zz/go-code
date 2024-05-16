package main

import (
	"math/bits"
	"slices"
)

/*
	方法二：线段树
	在方法一的暴力更新上优化，由于有区间的更新操作，需要使用 lazy 线段树。
	对于本题，在更新的时候需要优先递归右子树，从而保证是从右往左更新

	注：由于线段树常数比较大，在数据范围只有几百几千的情况下，不一定比方法一的暴力快。
*/

type seg []struct {
	l, r, cnt int
	todo      bool
}

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg) do(i int) {
	o := &t[i]
	o.cnt = o.r - o.l + 1
	o.todo = true
}

func (t seg) spread(o int) {
	if t[o].todo {
		t[o].todo = false
		t.do(o << 1)
		t.do(o<<1 | 1)
	}
}

// 查询区间 [l,r] o=1
func (t seg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].cnt
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return t.query(o<<1, l, r) + t.query(o<<1|1, l, r)
}

// 新增区间 [l,r] 后缀的 suffix 个时间点 o=1
// 相当于把线段树二分和线段树更新合并成了一个函数，时间复杂度为 O(log u)
func (t seg) update(o, l, r int, suffix *int) {
	size := t[o].r - t[o].l + 1
	if t[o].cnt == size { // 全部为运行中
		return
	}
	if l <= t[o].l && t[o].r <= r && size-t[o].cnt <= *suffix { // 整个区间全部该为运行中
		*suffix -= size - t[o].cnt
		t.do(o)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r > m { // 先更新右子树
		t.update(o<<1|1, l, r, suffix)
	}
	if *suffix > 0 { // 再更新左子树（如果还有需要新增的时间点）
		t.update(o<<1, l, r, suffix)
	}
	t[o].cnt = t[o<<1].cnt + t[o<<1|1].cnt
}

func findMinimumTime(tasks [][]int) int {
	slices.SortFunc(tasks, func(a, b []int) int { return a[1] - b[1] })
	u := tasks[len(tasks)-1][1]
	st := make(seg, 2<<bits.Len(uint(u-1)))
	st.build(1, 1, u)
	for _, t := range tasks {
		start, end, d := t[0], t[1], t[2]
		d -= st.query(1, start, end) // 去掉运行中的时间点
		if d > 0 {
			st.update(1, start, end, &d) // 新增时间点
		}
	}
	return st[1].cnt
}
