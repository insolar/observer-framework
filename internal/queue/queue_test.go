// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/observer-framework/blob/master/LICENSE.md.

// +build unit

package queue

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/insolar/observer-framework/internal"
)

func TestNewRecordQueue(t *testing.T) {
	timeout := time.Second
	q := New(timeout)
	require.NotNil(t, q)
	require.NotNil(t, q.entities)
	require.Equal(t, timeout, q.timeout)
}

func TestRecordQueue_Pop(t *testing.T) {
	q := New(time.Second)
	expected := &internal.RawRecord{
		RecordNumber: 100,
		Type:         "activate",
		PulseNumber:  100000,
	}
	var result *internal.RawRecord
	done := int32(0)
	timeout := int32(0)
	go func() {
		select {
		case q.entities <- expected:
			atomic.StoreInt32(&done, 1)
		case <-time.After(time.Second * 10):
			atomic.StoreInt32(&timeout, 1)
		}
	}()
	for atomic.LoadInt32(&(done)) == 0 && atomic.LoadInt32(&(timeout)) == 0 {
		r := q.Pop()
		if r != nil {
			result = r
		}
	}
	require.Equal(t, expected, result)
}

func TestRecordQueue_Pop_Empty(t *testing.T) {
	q := New(time.Second)
	r := q.Pop()
	require.Nil(t, r)
}

func TestRecordQueue_PopWithWaiting(t *testing.T) {
	q := New(time.Second * 10)
	expected := &internal.RawRecord{
		RecordNumber: 100,
		Type:         "activate",
		PulseNumber:  100000,
	}
	var r *internal.RawRecord
	done := make(chan struct{})
	go func() {
		r = q.PopWithWaiting()
		done <- struct{}{}
	}()
	q.entities <- expected
	<-done // wait for pop to return
	require.Equal(t, expected, r)
}

func TestRecordQueue_PopWithWaiting_Timeout(t *testing.T) {
	q := New(time.Millisecond)
	r := q.PopWithWaiting()
	require.Nil(t, r)
}

func TestRecordQueue_Push(t *testing.T) {
	q := New(time.Second)
	r := &internal.RawRecord{
		RecordNumber: 100,
		Type:         "activate",
		PulseNumber:  100000,
	}
	go func() {
		q.Push(r)
	}()
	expected := <-q.entities
	require.Equal(t, expected, r)
}
