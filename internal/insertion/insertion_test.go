package insertion

import (
	"math/rand"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	resChannel := make(chan bool)
	var wg sync.WaitGroup

	execSort := func(wg *sync.WaitGroup) {
		defer wg.Done()

		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		temp := r.Perm(100000)
		InsertionSort(temp)
		resChannel <- sort.IntsAreSorted(temp)
	}

	checkResult := func() {
		for res := range resChannel {
			assert.True(t, res)
		}
	}

	wg.Add(10)

	// start sorting asynchronously
	for n := 0; n < 10; n += 1 {
		go execSort(&wg)
	}

	// check result
	go checkResult()

	wg.Wait()
	defer close(resChannel)
}

func BenchmarkInsertionSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		InsertionSort(r.Perm(100000))
	}
}
