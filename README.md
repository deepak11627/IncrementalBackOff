# IncrementalBackOff
Incremental backoff implementation in Golang. This is a very simple backoff implementation which works as both exponential and then stabilies after certain period.

# Usage
It takes two params.
- First is the specify initialInterval that is to exponentially increase.
- Second param represent the duration after which the exponential growth of the backoff stops to stabilies itself to the constanct duration each time.
Eg:
```
backoff := NewIncrementalBackOff(1.25, 5)
for i := 1; i <= 10; i++ {
  fmt.Printf("Next Backoff is after %f seconds", backoff.NextBackOff())
}
```
# Output 
```
1.25, 1.56, 1.95, 2.44, 3.05, 3.81, 4.77, 5.00, 5.00, 5.00, 5.00
```
