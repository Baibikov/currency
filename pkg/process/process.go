package process

import (
	"context"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type Func func(ctx context.Context) error

type Process struct {
	ticker time.Duration
	mutex sync.Mutex
	entities []Func
}

func New(ticker time.Duration) *Process {
	return &Process{
		entities: make([]Func, 0),
		ticker: ticker,
	}
}

func (p *Process) Add(fn Func) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.entities = append(p.entities, fn)
}

func (p *Process) Wait(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			logrus.Info("process workers done")
			return
		case _, ok := <-time.Tick(p.ticker):
			if !ok {
				continue
			}

			for _, fn := range p.entities {
				err := fn(ctx)
				if err != nil {
					logrus.Errorf("proccess err %v", err)
					continue
				}

				logrus.Info("process worked")
			}
		}
	}
}