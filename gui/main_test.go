package gui

import (
	"testing"
)

func TestSliderValueIncreaseMayNotExceedMax(t *testing.T) {
	max := float32(2.0)
	got := getSliderValueFor(directionUp, 1, 0, 0.0, max)
	if got != max {
		t.Errorf("getSliderValueFor(inc,max=%f) = %f; want %f", max, got, max)
	}
}

func TestSliderValueDecreaseMayNotResultBelowMin(t *testing.T) {
	min := float32(25.0)
	got := getSliderValueFor(directionDown, 26, 0, min, 100)
	if got != min {
		t.Errorf("getSliderValueFor(dec,max=%f) = %f; want %f", min, got, min)
	}
}
