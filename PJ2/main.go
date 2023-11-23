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

	resultChannel := make(chan string)
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(3)

	go workerDomain.DoWork(3*time.Second, "Fetch Account", waitGroup, resultChannel)
	go workerDomain.DoWork(2*time.Second, "Get Account By Id", waitGroup, resultChannel)
	go workerDomain.DoWork(5*time.Second, "Search Account", waitGroup, resultChannel)

	go func() {
		waitGroup.Wait()
		close(resultChannel)
	}()

	// Lặp vô hạn để liên lục lấy dữ liệu từ channel, sẽ close khi channel bị close
	for res := range resultChannel {
		fmt.Printf("Result In Channel: %v \n", res)
	}

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
	domain.RunWorkCompose()
}
