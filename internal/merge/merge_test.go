package merge

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	for n := 0; n < 10; n++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		temp := r.Perm(100000)
		out := MergeSort(temp)
		assert.True(t, sort.IntsAreSorted(out))
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		MergeSort(r.Perm(100000))
	}
}
