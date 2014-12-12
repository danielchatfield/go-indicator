package indicator

import (
	"github.com/danielchatfield/go-chalk"
	"github.com/danielchatfield/go-log-symbols"
)

// Indicator represents the indicator state
type Indicator struct {
	symbol string
	index  int
}

// New returns a new Indicator
func New() *Indicator {
	s := &Indicator{
		symbol: chalk.Cyan(logsymbols.Spinner[0]),
	}

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
	s.symbol = chalk.Green(logsymbols.Success)
}

// Failure sets the failure symbol and sends a stop signal
func (s *Indicator) Failure() {
	s.symbol = logsymbols.Error
}
