package raft

import (
	"encoding/json"
	"io"
	"log"

	"github.com/hashicorp/raft"
)

type FSM struct {
	ctx *stCachedContext
	log *log.Logger
}

type logEntryData struct {
	Key   string
	Value string
}

// Apply applies a Raft Log entry to the key-value store.
func (f *FSM) Apply(logEntry *raft.Log) interface{} {
	e := logEntryData{}
	if err := json.Unmarshal(logEntry.Data, &e); err != nil {
		panic("Failed unmarshaling Raft Log entry. This is a bug.")
	}
	ret := f.ctx.StCached.CacheManager.Set(e.Key, e.Value)
	f.log.Printf("fms.Apply(), logEntry:%s, ret:%v\n", logEntry.Data, ret)
	return ret
}

// Snapshot returns a latest snapshot
func (f *FSM) Snapshot() (raft.FSMSnapshot, error) {
	return &snapshot{cm: f.ctx.StCached.CacheManager}, nil
}

// Restore stores the key-value store to a previous state.
func (f *FSM) Restore(serialized io.ReadCloser) error {
	return f.ctx.StCached.CacheManager.UnMarshal(serialized)
}
