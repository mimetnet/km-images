package main

import (
	"testing"
)

func TestDeviceBW(t *testing.T) {
	var d *Device

	if d = NewDevice("550"); d.Color {
		t.Error("550 is not a color device")
		t.FailNow()
	}
}

func TestDeviceColor(t *testing.T) {
	var d *Device

	if d = NewDevice("C550"); !d.Color {
		t.Error("C550 is a color device")
		t.FailNow()
	}

	if d = NewDevice("c550"); !d.Color {
		t.Error("c550 is a color device")
		t.FailNow()
	}
}
