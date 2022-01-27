// In this example we'll look at how to implement
// a _worker pool_ using goroutines and channels.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var que chan string

type dataS struct {
	Worker int `json:"worker"`
	Job    int `json:"job"`
}

func producer(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		// feed json data in the queue
		dd := dataS{
			Job:    j,
			Worker: id,
		}
		data, _ := json.Marshal(dd)
		que <- string(data)
		fmt.Println("worker", id, " job", j, "produced value----->", string(data))
	}

}

func consumer(id int) {
	for d := range que {
		fmt.Println("worker", id, "consuming  data------>", d)
	}

}

func createQueue(sizeOfPool int) {
	// if queue is not created then create
	if que == nil {
		que = make(chan string, sizeOfPool)
	}
	fmt.Println("created queue")
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()

	// In order to use our pool of workers we need to send
	// them work and collect their results. We make 2
	// channels for this.
	// we can make numJobs also configurable
	jobs := os.Getenv("jobs")
	numJobs, err := strconv.Atoi(jobs)
	if err != nil {
		fmt.Println("Please provide correct jobs size by using jobs=")
		os.Exit(1)
	}

	jobs1 := make(chan int, numJobs)
	que = make(chan string, numJobs)

	// This starts up 3 workers, initially blocked
	// because there are no jobs yet.
	poolSize := os.Getenv("pool")
	workingPool, err := strconv.Atoi(poolSize)
	if err != nil {
		fmt.Println("Please provide correct pool size by using pool= ")
		os.Exit(1)
	}
	var wg sync.WaitGroup
	wg.Add(workingPool)

	createQueue(workingPool)

	for w := 1; w <= workingPool; w++ {
		go producer(w, jobs1, &wg)
		go consumer(w)
	}

	for j := 1; j <= numJobs; j++ {
		jobs1 <- j
		time.Sleep(time.Second * 2)
	}
	close(jobs1)
	wg.Wait()
	close(que)
}
