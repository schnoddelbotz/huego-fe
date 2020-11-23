package gui

import (
	"testing"
)

func TestNewAppInitializesUI(t *testing.T) {
	app := newApp(nil, nil)
	if app.briChan == nil {
		t.Errorf("newApp(nil,nil) = %v; want non-nil", app)
	}
}
