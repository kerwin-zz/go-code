package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
WithValue函数能够将请求作用域的数据与 Context 对象建立关系。声明如下：
	func WithValue(parent Context, key, val interface{}) Context
WithValue返回父节点的副本，其中与key关联的值为val。
仅对API和进程间传递请求域的数据使用上下文值，而不是使用它来传递可选参数给函数。
所提供的键必须是可比较的，并且不应该是string类型或任何其他内置类型，以避免使用上下文在包之间发生冲
突。WithValue的用户应该为键定义自己的类型。为了避免在分配给interface{}时进行分配，上下文键通常
具有具体类型struct{}。或者，导出的上下文关键变量的静态类型应该是指针或接口。
*/

type TraceCode string

var wg2 sync.WaitGroup

func worker2(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string) // 在子goroutine中获取trace code
	if !ok {
		fmt.Println("invalid trace code")
	}
LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg2.Done()
}

func main() {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	// 在系统的入口中设置trace code传递给后续启动的goroutine实现日志数据聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "125123213213")
	wg2.Add(1)
	go worker2(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束
	wg2.Wait()
	fmt.Println("over")
}
