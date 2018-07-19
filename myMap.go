
package main

import "sync"

type myMap struct{
	Data map[string]Num
	Lock sync.RWMutex
}

func NewMap() *myMap{
	d := make(map[string]Num)
	return &myMap{Data: d}
} 

func (m myMap) SetMap(k string, v Num){
	m.Lock.Lock()
	defer m.Lock.Unlock()
	m.Data[k] = v
}

func (m myMap) GetMap(k string) (Num, bool){
	m.Lock.RLock()
	defer m.Lock.RUnlock()
	val, ok := m.Data[k]
	return val, ok 
}

func (m myMap) DeleteMap(k string){
	m.Lock.Lock()
	defer m.Lock.Unlock()
	delete(m.Data,k)
}
