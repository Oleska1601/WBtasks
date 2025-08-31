package main

import (
	"fmt"
	"sync"
)

// использование мапы и mutex

type Data struct {
	m  map[int]string
	mu sync.RWMutex
}

func NewData() *Data {
	return &Data{
		m: make(map[int]string),
	}
}

func (d *Data) Get(key int) (string, bool) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	v, ok := d.m[key]
	if !ok {
		return "", false
	}
	return v, true
}

func (d *Data) Put(key int, value string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.m[key] = value
}

func (d *Data) Delete(key int) {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.m, key)
}

func main() {
	data := NewData()
	wg := &sync.WaitGroup{}

	// конкуректная запись
	for i := range 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			data.Put(i, fmt.Sprintf("number %d", i))
		}(i)
	}

	// конкуретное чтение
	for i := range 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			v, ok := data.Get(i)
			if !ok {
				// значение не будет найдено, если какая-то горутина выше, отвечающая за запись по текущему ключу,
				// не успеет отработать до вызова горутины, отвечающей за чтение по этому же ключу
				fmt.Printf("no value found by key <%d>\n", i)
			} else {
				fmt.Printf("value <%s> found by key <%d>\n", v, i)
			}

		}(i)
	}

	wg.Wait()

}
