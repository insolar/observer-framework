// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/observer-framework/blob/master/LICENSE.md.

package queue

import (
	"time"

	"github.com/insolar/observer-framework/internal"
)

type Record struct {
	entities chan *internal.RawRecord
	timeout  time.Duration
}

func New(timeout time.Duration) *Record {
	return &Record{
		entities: make(chan *internal.RawRecord),
		timeout:  timeout,
	}
}

func (q *Record) Push(record *internal.RawRecord) {
	q.entities <- record
}

func (q *Record) Pop() *internal.RawRecord {
	select {
	case record := <-q.entities:
		return record
	default:
		return nil
	}
}

func (q *Record) PopWithWaiting() *internal.RawRecord {
	select {
	case record := <-q.entities:
		return record
	case <-time.After(q.timeout):
		return nil
	}
}

func (q *Record) Len() int {
	return 0
}
