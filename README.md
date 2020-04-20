# Go-TaskQueue-Lite 
[![GoDoc](https://godoc.org/github.com/sparrc/go-ping?status.svg)]()
![Go](https://github.com/ssubedir/go-taskqueue-lite/workflows/Go/badge.svg)

Lite Task/Job queue library using go rutines and channels 
  - Thread-Safe
  - Queue jobs/tasks 
  - Custom workers pool size

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
	"log"
	"os"
	tq "github.com/ssubedir/go-taskqueue-lite"
)
type TestTask struct {
	ID int
}
func (t *TestTask) Run() {
	fmt.Printf("I am working - '%d'\n", t.ID)
}
func main() {
	// Queue with 8 workers
	queue := tq.NewQueue(8)
	queue.Start()
	queue.Submit(&TestTask{1})
	queue.Submit(&TestTask{2})
	queue.Submit(&TestTask{3})
	queue.Stop()
}
```


Output
```sh
I am working - '1'
I am working - '2'
I am working - '3'
```


Task Interface:
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


License
----
Go-TaskQueue-Lite  is under [MIT](https://github.com/ssubedir/go-taskqueue-lite/blob/master/LICENSE) License.
See the LICENSE file for details.

