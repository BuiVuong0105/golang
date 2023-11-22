package main

import (
	"fmt"
	"pj2/domain"
	"sync"
	"time"
)

func goroutineExample() {

	start := time.Now()
	workerDomain := &domain.WokerDomain{}
	workerDomain.SetName("vuongbv")

	resultChannel := make(chan string)
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(3)
	// go workerDomain.DoWorkWaitGroup(3*time.Second, "Fetch Account", waitGroup)
	// go workerDomain.DoWorkWaitGroup(2*time.Second, "Get Account By Id", waitGroup)
	// waitGroup.Wait()

	go workerDomain.DoWork(3*time.Second, "Fetch Account", waitGroup, resultChannel)
	go workerDomain.DoWork(2*time.Second, "Get Account By Id", waitGroup, resultChannel)
	go workerDomain.DoWork(5*time.Second, "Search Account", waitGroup, resultChannel)

	// result1 := <-resultChannel
	// result2 := <-resultChannel
	// fmt.Printf("Result In Channel: result 1: %v, result 2: %v \n", result1, result2)

	go func() {
		for res := range resultChannel {
			fmt.Printf("Result In Channel: %v \n", res)
		}
		fmt.Printf("Total Dowork: %v second \n", time.Since(start))
	}()
	waitGroup.Wait()
	close(resultChannel)
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
	goroutineExample()
	// caculate()
	// anynomouseFunction()
}
