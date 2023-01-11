package safemap

import "testing"

// Normal Map is not a thread safe when you work with goroutine

func TestUnsafeMap(t *testing.T) {
	data := make(map[int]int)

	for i := 0; i < 10; i++ {
		go func(i int) {
			data[i] = i
		}(i)
	}
}

//	Safe Map is thread safe when you work with goroutine

func TestSafeMap(t *testing.T) {
	m := New[int, int]()

	for i := 0; i < 10000; i++ {
		go func(i int) {
			m.Insert(i, i*2)
			value, err := m.Get(i)
			if err != nil {
				t.Error(err)
			}
			if value != i*2 {
				t.Errorf("%d shoud not be %d", i, i*2)
			}
		}(i)
	}
}
