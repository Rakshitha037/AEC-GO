// package main
// import(
// 	"fmt"
// 	"sync"
// 	"time"

// )
// func worker(id int,checkpoint chan bool,resume chan bool,wg *sync.WaitGroup){
// 	defer wg.Done()
// 	fmt.Printf("worker %d:starting\n",id)
// 	time.Sleep(time.Duration(id) *time.Second)//simulate some work
// 	fmt.Printf("worker %d:Checkpoint reached\n",id)
// 	checkpoint<-true//signal that checkpoint has been reached
// 	<-resume//wait for the resume signal
// 	fmt.Printf("worker %d:Resuming\n",id)
// 	//continue with the remaining work
	

// }
// func main{
// 	numWorkers := 5
// 	checkpoint := make(chan bool)
// 	resume := make(chan bool)
// 	var wg sync.WaitGroup
// 	//launch the worker goroutine
// 	for i:=1;i<=numWorkers;i++{
// 		wg.Add(1)
// 		go worker(i,checkpoint,resume,&wg)
// 	}
// 	//wait for all the workers to reach checkpoits
// 	for i:=1;i<numWorkers;i++{
// 		<-checkpoint

// 	}
// 	fmt.Println("All workers reached the checkpoint")
// 	fmt.Println("Resuming all workers")
// 	//signal all workers to resume
// 	for i:=1;i<=numWorkers;i++{
// 		resume<-true
// 	}
// 	wg.Wait()
// 	fmt.Println("All workers completed their work")
// }
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, checkpoint chan bool, resume chan bool, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d:Starting\n", id)
    time.Sleep(time.Duration(id) * time.Second)
    fmt.Printf("Worker %d:Checkpoint reached\n", id)
    checkpoint <- true
    <-resume
    fmt.Printf("Worker %d:Resuming\n", id)
}
func main() {
    numWorkers := 5
    checkpoint := make(chan bool)
    resume := make(chan bool)
    var wg sync.WaitGroup
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go worker(i, checkpoint, resume, &wg)
    }
    for i := 1; i <= numWorkers; i++ {
        <-checkpoint
    }
    fmt.Println("All workers reached the checkpoint")
    fmt.Println("Resuming all workers")
    for i := 1; i <= numWorkers; i++ {
        resume <- true
    }
    wg.Wait()
    fmt.Println("All workers completed their work")

}

