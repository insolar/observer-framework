// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/observer-framework/blob/master/LICENSE.md.

// +build unit

package queue

import (
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
	go func() {
		q.entities <- expected
	}()
	time.Sleep(time.Millisecond)
	r := q.Pop()
	require.Equal(t, expected, r)
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
