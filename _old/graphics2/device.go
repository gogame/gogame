package graphics

import "github.com/gogame/gogame/graphics/sdl2"

type Device interface {
}

func NewDevice() (d Device, err error) {
	return sdl2.NewDevice()
}
