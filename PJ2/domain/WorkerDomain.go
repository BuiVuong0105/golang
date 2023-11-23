package domain

import (
	"fmt"
	"sync"
	"time"
)

type WokerDomain struct {
	name   string
	blance float64
}

// pointer receiver
// Chấp nhận cả WorkerDomain hoặc địa chỉ WorkerDomain, nếu truyền struct WorkerDomain thì sẽ tạo ra 1 con trỏ trỏ đến địa chỉ đó.
// Trong Go, khi bạn gọi một method với một pointer receiver trên một biến kiểu struct, Go sẽ tự động chuyển đổi biến thành một con trỏ tới biến đó (a pointer to the variable) nếu cần thiết.
// Điều này giúp người lập trình có thể gọi method với cả giá trị và con trỏ mà không cần phải quan tâm đến việc sử dụng pointer hay giá trị.

func (w *WokerDomain) SetName(name string) {
	w.name = name
}

func (w *WokerDomain) GetName() string {
	return w.name
}

func (w *WokerDomain) IncrementBlance(amount float64) float64 {
	w.blance = w.blance + amount
	return w.blance
}

func (w *WokerDomain) DoWork(duration time.Duration, workName string, waitGroup *sync.WaitGroup, channel chan string) {
	fmt.Printf("Begin Working: %s - %s \n", w.name, workName)
	time.Sleep(duration)
	fmt.Printf("Done Work: %s - %s \n", w.name, workName)
	channel <- workName
	waitGroup.Done()
}

func (w *WokerDomain) DoWorkChannel(duration time.Duration, workName string, channel chan string) {
	fmt.Printf("Begin Working: %s - %s \n", w.name, workName)
	time.Sleep(duration)
	fmt.Printf("Done Work: %s - %s \n", w.name, workName)
	channel <- workName
}
