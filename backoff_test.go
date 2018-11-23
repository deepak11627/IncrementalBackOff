package rabbitmq

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBackOff(t *testing.T) {
	backoff := NewIncrementalBackOff(1.25, 5)
	expectedBackOff := []float64{1.25, 1.56, 1.95, 2.44, 3.05, 3.81, 4.77, 5.00, 5.00, 5.00, 5.00}
	nextBackoff := backoff.NextBackOff()
	resultBackOff := []float64{nextBackoff}
	timer := time.NewTimer(time.Duration(nextBackoff) * time.Millisecond)
	counter := 1
	for {
		select {
		case <-timer.C:
			nextBackoff = backoff.NextBackOff()
			resultBackOff = append(resultBackOff, nextBackoff)
			counter++
		}
		if counter <= 10 {
			timer = time.NewTimer(time.Duration(nextBackoff) * time.Millisecond)
		} else {
			timer.Stop()
			break
		}
	}
	assert.Equal(t, expectedBackOff, resultBackOff)
}
