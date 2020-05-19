# Go-TaskQueue-Lite 
[![GoDoc](https://godoc.org/github.com/sparrc/go-ping?status.svg)](https://godoc.org/github.com/ssubedir/go-taskqueue-lite)
[![Go](https://github.com/ssubedir/go-taskqueue-lite/workflows/Go/badge.svg)](https://github.com/ssubedir/go-taskqueue-lite/actions)

An easy way to perform background processing in go. 
  - Thread-Safe
  - Custom workers pool size
  - Supported background tasks/jobs
  	- only once
  	- scheduled
  	- recurring
  

### Installation

Install Go-TaskQueue-Lite using go get

```sh
$ go get github.com/ssubedir/go-taskqueue-lite
```
# Example
Import this package and write
```go
package main
import (
	"fmt"
	tq "github.com/ssubedir/go-taskqueue-lite"
)
type TestTask struct {
	ID int
}
func (t *TestTask) Run() {
	fmt.Printf("Task - '%d'\n", t.ID)
}
func main() {
	// queue with 8 workers
	queue := tq.NewQueue(8)
	queue.Start()
	defer queue.Stop()
	// Run Tasks
	...
	...
	...
}
```
### Task Interface:
```go
// Task interface
type Task interface {
	Run()
}
```

All tasks must implement `Run()`

```go
type TestTask struct {
    // task struct
}
func (t *TestTask) Run() {
    // do task
}
```

## Only once task

Enqueue Parameters

```go
Enqueue(Task Task)
```

only once tasks are executed only once and almost immediately after creation. 
```go
queue := tq.NewQueue(8)
queue.Start()
defer queue.Stop()
queue.Enqueue(&TestTask{1})
```

## Delayed task

```go
Schedule(Task Task, duration_string  string)
```
A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as  "2h45m". Valid time units are "s", "m", "h". 

```

// example duration_string inputs

"10h"
"1h10m10s"
"1h1m"
"20s"
```


Delayed tasks are executed only once too, but not immediately, after a certain time interval.
```go
queue := NewQueue(8)
queue.Start()
defer queue.Stop()
queue.Schedule(&TestTask{1}, "10s")
```


## Recurring task

```go
Recurring(Task Task, duration_string  string)
```
A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as  "2h45m". Valid time units are "s", "m", "h". 

```
// example duration_string inputs

"10h"
"1h10m10s"
"1h1m"
"20s"
```

Recurring task are executed many times on the specified schedule.
```go
queue := NewQueue(8)
queue.Start()
defer queue.Stop()
queue.Recurring(&TestTask{1}, "1m10s")
```


### Todos
 - Batch tasks


## Built With

* [GO](https://golang.org/) - Programming language


## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/ssubedir/go-taskqueue-lite/blob/master/LICENSE) file for details

