// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/observer-framework/blob/master/LICENSE.md.

package internal

type RawRecord struct {
	RecordNumber        uint32
	Reference           []byte
	Type                string // enum of request, result, state
	ObjectReference     []byte
	PrototypeReference  []byte
	Payload             []byte
	PrevRecordReference []byte
	PulseNumber         uint32
	Timestamp           uint32
}

type RecordQueue interface {
	Pop() *RawRecord
	Push(record *RawRecord)
	PopWithWaiting() *RawRecord
	Len() int
}
