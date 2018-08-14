package mocking

import (
	"fmt"
	"io"
)

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

//Countdown the no's
func Countdown(out io.Writer, spy *CountdownOperationsSpy) {
	for i := 3; i > 0; i-- {
		spy.Sleep()
		fmt.Fprintln(out, i)
	}
	spy.Sleep()
	fmt.Fprint(out, "Go!")
}
