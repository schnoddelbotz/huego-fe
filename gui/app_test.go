package gui

import (
	"testing"
)

func TestNewAppInitializesUI(t *testing.T) {
	app := newApp(nil, nil, "")
	if app.ctrlChan == nil {
		t.Errorf("newApp(nil,nil) = %v; want non-nil", app)
	}
}
