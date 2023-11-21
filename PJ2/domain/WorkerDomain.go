package domain

import (
	"fmt"
	"sync"
	"time"
)

type WokerDomain struct {
	Name string
}

func (w *WokerDomain) DoWork(duration time.Duration, workName string, waitGroup *sync.WaitGroup, channel chan string) {
	fmt.Printf("Begin Working: %s - %s \n", w.Name, workName)
	time.Sleep(duration)
	fmt.Printf("Done Work: %s - %s \n", w.Name, workName)
	channel <- workName
	waitGroup.Done()
}

func (w *WokerDomain) DoWorkChannel(duration time.Duration, workName string, channel chan string) {
	fmt.Printf("Begin Working: %s - %s \n", w.Name, workName)
	time.Sleep(duration)
	fmt.Printf("Done Work: %s - %s \n", w.Name, workName)
	channel <- workName
}
