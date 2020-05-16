package job

import (
	"log"

	"github.com/manue1/myhttp/internal/result"
)

type (
	// Output is used to process the results
	Output struct {
		outputPrinter
	}
	outputPrinter interface {
		Print(chan struct{}, chan result.Page)
	}
)

// NewOutput returns a new default Output
func NewOutput() Output {
	return Output{}
}

// Print is the default output method which logs the results on stdout
func (o Output) Print(done chan struct{}, results chan result.Page) {
	for r := range results {
		log.Print(r.String())
	}

	done <- struct{}{}
}
