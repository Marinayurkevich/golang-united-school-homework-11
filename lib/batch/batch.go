package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	ch := make(chan int, pool)
	var wg sync.WaitGroup
	var mx sync.Mutex
	var i int64
	for i = 0; i < n; i++ {
		wg.Add(1)
		ch <- 0
		go func(j int64) {
			OneUser := getOne(j)
			mx.Lock()
			res = append(res, OneUser)
			mx.Unlock()
			wg.Done()
			<-ch
		}(i)
	}
	wg.Wait()
	return res
}
