package integers

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
}

type CountdownOperationsSpy struct {
	Calls []string
}

const write = "write"
const sleep = "sleep"

func Countdown(out io.Writer, sleeper *CountdownOperationsSpy) {
	for i := 3; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintf(out, "%d\n", i)
	}

	sleeper.Sleep()
	fmt.Fprintf(out, "%s", "Go!")
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}
