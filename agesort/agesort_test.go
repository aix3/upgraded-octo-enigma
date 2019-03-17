package agesort

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"testing"
)

func TestAgeSort(t *testing.T) {
	max := 100

	ages1 := make([]int, 10000)
	ages2 := make([]int, 10000)

	for i := 0; i < 10000; i++ {
		r := rand.Intn(max)
		ages1[i] = r
		ages2[i] = r
	}

	sort.Ints(ages1)
	agesort(ages2, max)

	if !strings.EqualFold(fmt.Sprintf("%v", ages1), fmt.Sprintf("%v", ages2)) {
		t.Errorf("Expected: %v...%v, actually: %v...%v", ages1[0:10], ages1[len(ages1)-10:len(ages1)], ages2[0:10], ages2[len(ages2)-10:len(ages2)])
	}
}
