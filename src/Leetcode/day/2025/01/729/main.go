package main

import "github.com/emirpasic/gods/v2/trees/redblacktree"

/*
	实现一个 MyCalendar 类来存放你的日程安排。如果要添加的日程安排不会造成 重复预订 ，则可以存储这个新的日程安排。

	当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生 重复预订 。

	日程可以用一对整数 startTime 和 endTime 表示，这里的时间是半开区间，即 [startTime, endTime), 实数 x 的范围为，  startTime <= x < endTime 。

	实现 MyCalendar 类：
	· MyCalendar() 初始化日历对象。
	· boolean book(int startTime, int endTime) 如果可以将日程安排成功添加到日历中而不会导致重复预订，返回 true 。否则，返回 false 并且不要将该日程安排添加到日历中。
*/

type MyCalendar struct {
	*redblacktree.Tree[int, int]
}

func Constructor() MyCalendar {
	return MyCalendar{redblacktree.New[int, int]()}
}

func (c *MyCalendar) Book(startTime int, endTime int) bool {
	if p, ok := c.Ceiling(startTime + 1); ok && p.Value < endTime {
		return false
	}
	c.Put(endTime, startTime)
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */
