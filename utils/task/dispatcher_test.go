package task

import (
	"testing"

	"github.com/hellodudu/Ultimate/iface"
	logger "github.com/hellodudu/Ultimate/utils/log"
)

var chCallback chan int

func init() {
	logger.Init(false, false, "dispatcher_test")
	chCallback = make(chan int, 1)
}

func taskCallback(_ iface.ITCPConn, data []byte) {
	str := string(data[:])

	if str != "hello world" {
		chCallback <- 0
		return
	}

	chCallback <- 1
}

func TestNewDispatcher(t *testing.T) {

	td, err := NewDispatcher()
	if err != nil {
		t.Error("NewDispatcher failed")
	}

	td.AddTask(&TaskReqInfo{
		Data: []byte("hello world"),
		CB:   taskCallback,
	})

	success := <-chCallback
	if success != 1 {
		t.Error("AddTask failed")
	}

	close(chCallback)
}
