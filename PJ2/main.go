package main

import (
	"fmt"
	"pj2/domain"
	"sync"
	"time"
)

func goroutineExample() {

	start := time.Now()
	workerDomain := domain.WokerDomain{}
	workerDomain.SetName("vuongbv")

	// Sử dụng channel thì việc gửi và nhận phải ở trong các gorountine khác nhau, khi gửi vào thì nếu vượt quá buffer mà channel không có ai đăng ký nhận sẽ bị báo lỗi
	resultChannel := make(chan string, 10)
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(3)

	go workerDomain.DoWork(1*time.Second, "Fetch Account", waitGroup, resultChannel)
	go workerDomain.DoWork(2*time.Second, "Get Account By Id", waitGroup, resultChannel)
	go workerDomain.DoWork(5*time.Second, "Search Account", waitGroup, resultChannel)

	go func() {
		waitGroup.Wait()
		close(resultChannel)
	}()

	workerDomain.ReadChannel(resultChannel)

	fmt.Printf("Total Dowork: %v second \n", time.Since(start))
}

// func divide(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("Cannot divide by zero")
// 	}
// 	return (a / b), nil
// }

// func caculate() {
// 	result, err := divide(5, 0)
// 	if err != nil {
// 		fmt.Println("ERROR: ", err)
// 	} else {
// 		fmt.Println("RESULT: ", result)
// 	}
// }

// func anynomouseFunction() {
// 	for i := 0; i <= 10; i++ {
// 		func(i int) {
// 			fmt.Println("Anynomouse Function: ", i)
// 		}(i)
// 	}
// }

func main() {
	// goroutineExample()
	// caculate()
	// anynomouseFunction()
	// domain.RunWriter()
	// domain.RunIncrement()
	// domain.RunCompose()
	// domain.RunWorkCompose()
	// domain.RunInterface()

	wg := sync.WaitGroup{}
	wg.Add(2)
	resultChannel := make(chan int, 2)

	go func() {
		fmt.Println("RS: ", <-resultChannel) // Đoc channel nếu đọc mà không có data thì báo lỗi
		wg.Done()
	}()

	go func() {
		resultChannel <- 42 // khi gửi vào thì nếu vượt quá buffer mà channel không có ai đăng ký nhận sẽ bị báo lỗi
		wg.Done()
	}()
	wg.Wait()
}
