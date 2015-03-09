package barrier

import "sync"

type Barrier struct{
    lock  sync.Mutex
    cond  sync.Cond
    threshold  int
    count      int
    cycle      bool
}

func NewBarrier(n  int) *Barrier{
    b := &Barrier{threshold: n, count: n}
    b.cond.L = &b.lock
    return b
}

func (b *Barrier) Wait() (last bool){
    b.lock.Lock()
    defer  b.lock.Unlock()
    cycle :=  b.cycle
    b.count--
    if b.count == 0 {
       b.cycle  =  !b.cycle 
       b.count = b.threshold 
       b.cond.Broadcast()
       last = true
    }else{
      for cycle == b.cycle {
          b.cond.Wait()
      }
    }
    return
}
