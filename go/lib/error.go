package lib

import "fmt"

var QueueInputIsNil = fmt.Errorf("input to queue creation can't be nil")
var QueueIsEmpty = fmt.Errorf("queue is empty")

var StackInputIsNil = fmt.Errorf("input to stack creation can't be nil")
var StackIsEmpty = fmt.Errorf("stack is empty")
