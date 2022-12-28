package main 

import (
	"fmt"
	"sync"
	"time"
)
type Loader struct {
	finalCapacity int64
	current int64
	inProgress bool  
	wg sync.WaitGroup
	waiting bool
}

func (l *Loader) load() {
	ch1 := make(chan bool)
	go l.print(ch1)
	<-ch1
}

func (l *Loader) print(ch1 chan bool) {
	var current int64 = 0 
	//fmt.Println(l.current, "/", l.finalCapacity)
	l.inProgress = true 
	//l.wg.Add(1)
	for ; current < l.finalCapacity; current = l.current  {
		l.wg.Add(1)
		l.waiting = true 
		fmt.Print("-") 
		l.wg.Wait()
		l.waiting = false 
	} 
	fmt.Println()
	fmt.Println("Done")
	l.inProgress = false 
	ch1 <- true 

}

func (l *Loader) increment(byVal int64) {
	l.current = l.current + byVal 
	if l.waiting {
		l.wg.Done()
	}
}
func main() {
	l := Loader{finalCapacity: 20}	
	
	stopCh := make(chan bool)
	go l.load()
	go func(){
		const BYVAL int64 = 2
		var i int64 = 0
		for ; i <= (l.finalCapacity / BYVAL); i+= 1 {
			time.Sleep(300 * time.Millisecond)
			l.increment(BYVAL)
			
			//fmt.Println("incrementer", l.current)
			//fmt.Print("incremented")
			
		}
		stopCh <- true 
	}()
	<-stopCh 
}