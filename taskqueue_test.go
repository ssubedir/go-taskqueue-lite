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
