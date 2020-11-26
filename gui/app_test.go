package gui

import (
	"testing"
)

func TestNewAppInitializesUI(t *testing.T) {
	app := newApp(nil, nil, "", "")

	if app.ctrlChan == nil {
		t.Errorf("newApp(nil,nil, ...) = %v; want non-nil", app)
	}
}

func TestLightFilterInitialization(t *testing.T) {
	lightFilter := "1,2,hello,world,3"
	expectedLightFilter := []int{1, 2, 3}

	app := newApp(nil, nil, lightFilter, "")

	if len(app.lightFilter) != 3 {
		t.Errorf(`newApp(nil,nil,"%s",...) = %d; want 3`, lightFilter, len(app.lightFilter))
	}
	for idx, expected := range expectedLightFilter {
		if app.lightFilter[idx] != expected {
			t.Errorf(`newApp(nil,nil,...,"%s") [%d] = %d; want %d`, lightFilter, idx, app.groupFilter[idx], expected)
		}
	}
}

func TestGroupFilterInitialization(t *testing.T) {
	groupFilter := "4,,56,oh no,2,9"
	expectedGroupFilter := []int{4, 56, 2, 9}

	app := newApp(nil, nil, "", groupFilter)

	if len(app.groupFilter) != 4 {
		t.Errorf(`newApp(nil,nil,...,"%s") = %d; want 3`, groupFilter, len(app.groupFilter))
	}
	for idx, expected := range expectedGroupFilter {
		if app.groupFilter[idx] != expected {
			t.Errorf(`newApp(nil,nil,...,"%s") [%d] = %d; want %d`, groupFilter, idx, app.groupFilter[idx], expected)
		}
	}
}
