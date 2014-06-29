// A simple semaphore primitive
//
// Usage:
//    sem := New(3)
//    for i := 0; i < 10; i++ {
//    	sem.Wait()
//    	go func() {
//    		time.Sleep(time.Second)
//    		sem.Signal()
//    	}()
//    }
package semaphore

type Semaphore chan struct{}

// Increments the `Semaphore` by 1
func (s Semaphore) Signal() {
	s <- struct{}{}
}


// Decrements the `Semaphore` by 1
func (s Semaphore) Wait() {
	<-s
}


// New returns an instance of a `Semaphore`:
//    sem := New(10)
func New(size int) Semaphore {
	sem := make(Semaphore, size)
	for _ = range make([]struct{}, size) {
		sem <- struct{}{}
	}
	return sem
}
