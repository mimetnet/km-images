package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Device struct {
	Pro   bool
	Press bool
	Color bool
	Name  string
}

func NewDevice(name string) *Device {
	d := new(Device)
	d.Color = false

	if 'C' == name[0] {
		name = strings.TrimLeft(name, "C")

		d.Name = name
		d.Color = true

		if strings.HasSuffix(name, "L") {
			d.Pro = true
		}
	} else {
		d.Name = name
	}

	if !d.Pro && !d.Press {
		size, _ := strconv.ParseUint(name, 10, 32)

		if 999 < size {
			d.Press = true
		} else if 900 < size {
			d.Pro = true
		}
	}

	return d
}

func (d *Device) String() string {
	if d.Color {
		if d.Pro {
			return fmt.Sprintf("Pro C%s", d.Name)
		} else if d.Press {
			return fmt.Sprintf("Press C%s", d.Name)
		} else {
			return fmt.Sprintf("C%s", d.Name)
		}
	}

	if d.Pro {
		return "Pro " + d.Name
	} else if d.Press {
		return "Press " + d.Name
	} else {
		return d.Name
	}
}
