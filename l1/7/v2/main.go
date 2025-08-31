package main

import (
	"fmt"
	"sync"
)

// использование sync.Map

func main() {
	m := sync.Map{}
	wg := &sync.WaitGroup{}

	// конкуректная запись
	for i := range 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Store(i, fmt.Sprintf("number %d", i))
		}(i)
	}

	// конкуретное чтение
	for i := range 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			v, ok := m.Load(i)
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
