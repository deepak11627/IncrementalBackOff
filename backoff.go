package main

import (
	"fmt"
	"math"
	"strconv"
)

type BackOff interface {
	NextBackOff() float64
	Reset()
}
type IncrementalBackOff struct {
	initialInterval float64
	stableInterval  float64
	counter         int
}

// NewIncrementalBackOff takes time in seconds and returns an IncrementalBackOff
func NewIncrementalBackOff(interval, stableInterval float64) *IncrementalBackOff {
	b := &IncrementalBackOff{
		initialInterval: interval,
		stableInterval:  stableInterval,
		counter:         1,
	}
	b.Reset()
	return b
}

// Reset the BackOff
func (b *IncrementalBackOff) Reset() {
	b.counter = 1
}

// NextBackOff return the duration for the NextBackOff
func (b *IncrementalBackOff) NextBackOff() float64 {
	backoff := math.Pow(b.initialInterval, float64(b.counter))
	b.counter++
	if backoff >= b.stableInterval {
		backoff = b.stableInterval
	}
	backoff, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", backoff), 64)
	return backoff
}
