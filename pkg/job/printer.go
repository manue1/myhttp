package job

import (
	"log"

	"github.com/manue1/myhttp/internal/result"
)

type (
	Output struct {
		outputPrinter
	}
	outputPrinter interface {
		Print(chan struct{}, chan result.Page)
	}
)

func NewOutput() Output {
	return Output{}
}

func (o Output) Print(done chan struct{}, results chan result.Page) {
	for r := range results {
		log.Print(r.String())
	}

	done <- struct{}{}
}
