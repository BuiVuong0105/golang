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
