package main

import (
	"sync"
	"time"
	"fmt"
)

type CompletionManager struct {
	mu sync.Mutex
	doneCh map[string]chan struct{}
}

func NewCompletionManager() *CompletionManager {
	return &CompletionManager{
		doneCh: make(map[string]chan struct{}),
	}
}

func (cm *CompletionManager) Wait(uuid string) {
	cm.mu.Lock()
	ch, ok := cm.doneCh[uuid]
	if !ok {
		ch = make(chan struct{}, 1)
		cm.doneCh[uuid] = ch
	}
	cm.mu.Unlock()

	select {
	case <-ch:
		return
	case <-time.After(10 * time.Second):
		fmt.Println("Timed out waiting for completion on characteristic:", uuid)
	}
}

func (cm *CompletionManager) Notify(uuid string) {
	cm.mu.Lock()
	ch, ok := cm.doneCh[uuid]
	if !ok {
		ch = make(chan struct{}, 1)
		cm.doneCh[uuid] = ch
	}
	cm.mu.Unlock()

	select {
	case ch <- struct{}{}:
	default:
	}
}
