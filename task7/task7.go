package main

import (
	"fmt"
	"strconv"
	"sync"
)

type any = interface{}

type SyncMap struct {
	mx *sync.RWMutex
	mp map[any]interface{}
}

func NewSyncMap() *SyncMap{
	return &SyncMap{
		mp: make(map[any]interface{}),
		mx: &sync.RWMutex{},
	}
}

func (s *SyncMap) Load(key any) (value any, ok bool) {
	s.mx.RLock()
	defer s.mx.RUnlock()
	value, ok = s.mp[key]
	return
}

func (s *SyncMap) Store(key any, value any) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.mp[key] = value
}

func (s *SyncMap) Delete(key any) {
	s.mx.Lock()
	defer s.mx.Unlock()
	delete(s.mp, key)
}

func main() {
	sm := NewSyncMap()

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			defer wg.Done()
			sm.Store(i, strconv.Itoa(i) + "-v")
		}(&wg, i)
	}
	wg.Wait()

	for i := 0; i < 5; i++ {
		fmt.Println(sm.Load(i))
	}
}
