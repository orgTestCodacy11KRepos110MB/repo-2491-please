package core

import (
	"sync/atomic"
)

// An atomicBool is a type we use as an opaque boolean that doesn't trigger the race detector.
type atomicBool struct {
	b int32
}

func (b *atomicBool) SetTrue() {
	atomic.StoreInt32(&b.b, 1)
}

func (b *atomicBool) SetFalse() {
	atomic.StoreInt32(&b.b, 0)
}

func (b *atomicBool) Set(val bool) {
	if val {
		b.SetTrue()
		return
	}
	b.SetFalse()
}

func (b *atomicBool) Value() bool {
	return atomic.LoadInt32(&b.b) == 1
}
