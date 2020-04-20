// Go-TaskQueue-Lite

// Lite Task/Job queue library using go rutines and channels

//    Thread-Safe
//    Queue jobs/tasks
//    Custom workers pool size

// Installation

// Install Go-TaskQueue-Lite using go get

// $ go get github.com/ssubedir/go-taskqueue-lite

// Example

// Import this package and write

// package main
// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	tq "github.com/ssubedir/go-taskqueue-lite"
// )
// type TestTask struct {
// 	ID int
// }
// func (t *TestTask) Run() {
// 	fmt.Printf("I am working - '%d'\n", t.ID)
// }
// func main() {
// 	// Queue with 8 workers
// 	queue := tq.NewQueue(8)
// 	queue.Start()
// 	queue.Submit(&TestTask{1})
// 	queue.Submit(&TestTask{2})
// 	queue.Submit(&TestTask{3})
// 	queue.Stop()
// }

// Output

// I am working - '1'
// I am working - '2'
// I am working - '3'

// Task Interface:

// // Task interface
// type Task interface {
// 	Run()
// }

// All tasks must implement Run()

// type TestTask struct {
//     // task struct
// }
// func (t *TestTask) Run() {
//     // do task
// }


package taskqueue

import (
	"fmt"
	"testing"
)

// TestTask - holds only an ID to show state
type TestTask struct {
	ID int
}

// Process - test process function
func (t *TestTask) Run() {
	fmt.Printf("Processing Task '%d'\n", t.ID)
}

func TestQueue(t *testing.T) {
	queue := NewQueue(100)
	queue.Start()
	queue.Submit(&TestTask{1})
	queue.Submit(&TestTask{2})
	queue.Submit(&TestTask{3})
	queue.Stop()
}
