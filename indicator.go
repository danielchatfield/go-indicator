package indicator

import (
	"github.com/danielchatfield/go-chalk"
	"github.com/danielchatfield/go-log-symbols"
)

// Indicator represents the indicator state
type Indicator struct {
	symbol string
	index  int
	stop   chan struct{}
}

// New returns a new Indicator
func New() *Indicator {
	s := &Indicator{
		symbol: chalk.Cyan(logsymbols.Spinner[0]),
		stop:   make(chan struct{}),
	}

	go func() {
		for {
			select {
			case <-s.stop:
				return
			}
		}
	}()

	return s
}

func (s *Indicator) String() string {
	return string(s.symbol)
}

// Next increments the spinner frame index
func (s *Indicator) Next() *Indicator {
	s.index++
	s.symbol = chalk.Cyan(logsymbols.Spinner[s.index%len(logsymbols.Spinner)])
	return s
}

// Success sets the success symbol and sends a stop signal
func (s *Indicator) Success() {
	close(s.stop)
	s.symbol = logsymbols.Success
}

// Failure sets the failure symbol and sends a stop signal
func (s *Indicator) Failure() {
	close(s.stop)
	s.symbol = logsymbols.Error
}
