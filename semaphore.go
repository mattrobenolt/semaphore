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

type semaphore chan struct{}

// Increments the semaphore by 1
func (s semaphore) Signal() {
	s <- struct{}{}
}


// Decrements the semaphore by 1
func (s semaphore) Wait() {
	<-s
}


// New returns an instance of a semaphore:
//    sem := New(10)
func New(size int) semaphore {
	sem := make(semaphore, size)
	for i := 0; i < size; i++ {
		sem <- struct{}{}
	}
	return sem
}
