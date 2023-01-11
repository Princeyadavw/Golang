//  A worker and work queue example 

package main
 
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
 
var (
	
	wg sync.WaitGroup
	
	jobChan = make(chan int, 11)
	
	waiters = []int{0, 1, 2, 3, 4}
)
 
func main() {
	rand.Seed(time.Now().UnixNano())
 
	fmt.Println("BEGIN")
 
	
	wg.Add(1)
 
	
	go worker(jobChan, &wg)
 
	
	for i := 1; i <= 10; i++ {
		if !queueJob(i, jobChan) {
			fmt.Println("Channel is full... Service unavailable...")
		}
	}
 
	
	close(jobChan)
 
	// Block exiting until all the goroutines are finished.
	wg.Wait()
 
	fmt.Println("END")
}
 

func queueJob(job int, jobChan chan<- int) bool {
	select {
	case jobChan <- job:
		return true
	default:
		return false
	}
}
 
	
func worker(jobChan <-chan int, wg *sync.WaitGroup) {
	
	defer wg.Done()
 
	fmt.Println("Worker is waiting for jobs")
 
	for job := range jobChan {
		fmt.Println("Worker picked job", job)
 
		wait := time.Duration(rand.Intn(len(waiters)))
		time.Sleep(wait * time.Second)
 
		
		fmt.Println("Worker completed job", job, "in", int(wait), "second(s)")
	}
}