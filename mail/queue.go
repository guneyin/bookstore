package mail

import (
	"context"
)

type queue struct {
	jobs   chan job
	ctx    context.Context
	cancel context.CancelFunc
}

type job struct {
	Composer *Composer
	Action   func(c *Composer) error
}

func newQueue() *queue {
	ctx, cancel := context.WithCancel(context.Background())

	return &queue{
		jobs:   make(chan job),
		ctx:    ctx,
		cancel: cancel,
	}
}

func (q *queue) addJob(c *Composer, act func(c *Composer) error) {
	q.jobs <- job{Composer: c, Action: act}
}

func (j job) run() error {
	return j.Action(j.Composer)
}
