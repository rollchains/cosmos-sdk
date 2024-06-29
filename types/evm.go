package types

import (
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// used in context.go
type EVMEventManager struct {
	events []*ethtypes.Log
}

func NewEVMEventManager() *EVMEventManager {
	return &EVMEventManager{events: []*ethtypes.Log{}}
}

func (eem *EVMEventManager) Events() []*ethtypes.Log { return eem.events }

func (eem *EVMEventManager) EmitEvents(events []*ethtypes.Log) {
	eem.events = append(eem.events, events...)
}
