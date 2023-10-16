package insertion

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	for n := 0; n < 10; n++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		temp := r.Perm(100000)
		InsertionSort(temp)
		assert.True(t, sort.IntsAreSorted(temp))
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		InsertionSort(r.Perm(100000))
	}
}
