package memo_test

import (
	"golearn/ch9/memo4"
	"golearn/ch9/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}

/*
$ go test -v golearn/ch9/memo4
$ go test -run=TestConcurrent -race -v golearn/ch9/memo4
*/
