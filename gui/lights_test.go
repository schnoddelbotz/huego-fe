package gui

import (
	"testing"
)

func TestGetSliceIndex(t *testing.T) {
	haystack := []int{1, 2, 3, 4}
	needle := 3
	expectedIndex := 2
	result := sliceIndex(haystack, needle)
	if result != expectedIndex {
		t.Errorf("TestGetSliceIndex(%v,%d) = %d; want %d", haystack, needle, result, expectedIndex)
	}
}

func TestGetSliceIndexNegativeIfNotFound(t *testing.T) {
	haystack := []int{1, 2, 3, 4}
	needle := 5
	expectedIndex := -1
	result := sliceIndex(haystack, needle)
	if result != expectedIndex {
		t.Errorf("TestGetSliceIndex(%v,%d) = %d; want %d", haystack, needle, result, expectedIndex)
	}
}

func TestGetLightIDHigherThan(t *testing.T) {
	haystack := []int{2, 5, 9, 25}
	expectedID := 9
	result := idHigherThan(5, haystack)
	if result != expectedID {
		t.Errorf("TestGetLightIDHigherThan(5,%v) = %d; want %d", haystack, result, expectedID)
	}
}

func TestGetLightIDHigherThanSingleLight(t *testing.T) {
	haystack := []int{2}
	expectedID := 2
	result := idHigherThan(2, haystack)
	if result != expectedID {
		t.Errorf("TestGetLightIDHigherThan(2,%v) = %d; want %d", haystack, result, expectedID)
	}
}

func TestGetLightIDHigherThanAtUpperBound(t *testing.T) {
	haystack := []int{2, 5, 9, 25}
	expectedID := 25
	result := idHigherThan(25, haystack)
	if result != expectedID {
		t.Errorf("TestGetLightIDHigherThan(5,%v) = %d; want %d", haystack, result, expectedID)
	}
}

func TestGetLightIDLowerThan(t *testing.T) {
	haystack := []int{2, 5, 9, 25}
	expectedID := 2
	result := idLowerThan(5, haystack)
	if result != expectedID {
		t.Errorf("TestGetLightIDLowerThan(5,%v) = %d; want %d", haystack, result, expectedID)
	}
}

func TestGetLightIDLowerThanSingleLight(t *testing.T) {
	haystack := []int{2}
	expectedID := 2
	result := idLowerThan(2, haystack)
	if result != expectedID {
		t.Errorf("TestGetLightIDLowerThan(2,%v) = %d; want %d", haystack, result, expectedID)
	}
}

func TestGetLightIDLowerThanAtLowerBound(t *testing.T) {
	haystack := []int{2, 5, 9, 25}
	expectedID := 2
	result := idLowerThan(2, haystack)
	if result != expectedID {
		t.Errorf("TestGetLightIDLowerThan(2,%v) = %d; want %d", haystack, result, expectedID)
	}
}
