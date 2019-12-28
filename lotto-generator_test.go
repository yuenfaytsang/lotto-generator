package main

import (
	"sort"
	"testing"
)

func TestNumLength(t *testing.T) {
	if len(lotteryNumbers(6, 49)) != 6 {
		t.Errorf("Failed testing []int length for 6of49")
	}

	if len(lotteryNumbers(5, 50)) != 5 {
		t.Errorf("Failed testing []int length for Euro Lotto")
	}
}

func TestSortedNums(t *testing.T) {
	if sort.IntsAreSorted(lotteryNumbers(6, 49)) == false {
		t.Errorf("Failed testing sorted []int. Slice is not sorted")
	}
	if sort.IntsAreSorted(lotteryNumbers(5, 50)) == false {
		t.Errorf("Failed testing sorted []int. Slice is not sorted")
	}
}

func TestUniqueNum(t *testing.T) {
	a := generateNumber(100)
	b := generateNumber(100)

	if a == b {
		t.Errorf("Failed testing random num. Number is not Random")
	}
}
