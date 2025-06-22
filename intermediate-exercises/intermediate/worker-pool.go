package intermediate

import (
	"fmt"
	"sync"
	"time"
)

func doWork(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Millisecond * 100)
		results <- j * j
		fmt.Println("worker", id, "finished job", j)
	}
}

func ExecuteWorkerPool() {
	const numJobs = 100
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup
	for w := 1; w <= 10; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			doWork(id, jobs, results)
		}(w)
	}

	for i := range numJobs {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
	close(results)
	for a := 1; a <= numJobs; a++ {
		res := <-results
		fmt.Println("Results:", res)
	}
}
