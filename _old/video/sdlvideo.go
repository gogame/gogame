//+build !js
package video

import (
	"github.com/veandco/go-sdl2/sdl"
)

var ActiveDevice = &SdlDevice{}

type SdlDevice struct {
	window  *sdl.Window
	surface *sdl.Surface
}

func init() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
}

func (g *SdlDevice) Start() {
	var err error
	//mayby hide window to allow size change in user start displaying at first update of after start
	g.window, err = sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	g.surface, err = g.window.GetSurface()
	if err != nil {
		panic(err)
	}

}

func (g *SdlDevice) Update() {
	rect := sdl.Rect{0, 0, 200, 200}
	g.surface.FillRect(&rect, 0xffff0000)
	g.window.UpdateSurface()
}

func (g *SdlDevice) Exit() {
	g.window.Destroy()
	sdl.Quit()
}
