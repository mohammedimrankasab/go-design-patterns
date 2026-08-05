package singleton_test

import (
	"sync"
	"testing"

	"github.com/mohammedimrankasab/go-design-patterns/creational/singleton"
)

func TestSingleton_ReturnsSameInstance(t *testing.T) {

	first := singleton.GetInstance()
	second := singleton.GetInstance()

	if first != second {
		t.Fatal("expected same singleton instance")
	}
}

func TestSingleton_SetAndGet(t *testing.T) {

	config := singleton.GetInstance()

	config.Set(
		"database",
		"postgres",
	)

	value, exists := config.Get(
		"database",
	)

	if !exists {
		t.Fatal("expected database configuration")
	}

	if value != "postgres" {
		t.Fatalf(
			"expected postgres, got %s",
			value,
		)
	}
}

func TestSingleton_ConcurrentAccess(t *testing.T) {

	const workers = 100

	instances := make(
		[]*singleton.ConfigManager,
		workers,
	)

	var wg sync.WaitGroup

	wg.Add(workers)

	for i := 0; i < workers; i++ {

		go func(index int) {

			defer wg.Done()

			instances[index] = singleton.GetInstance()

		}(i)
	}

	wg.Wait()

	first := instances[0]

	for _, instance := range instances {

		if instance != first {
			t.Fatal(
				"expected all goroutines to receive same instance",
			)
		}
	}
}
