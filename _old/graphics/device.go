package graphics

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Device struct {
	window *sdl.Window
	render *sdl.Renderer
	//	title string
}

type Config struct {
	Title      string
	Width      int
	Height     int
	Scale      int
	Fullscreen bool
}

func init() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
}

//todo delete this function
func (d *Device) Render() *sdl.Renderer {
	return d.render
}

//func Title(t string) (*Device) {return }
//todo: implement fullscreen
func NewDevice(c *Config) (d *Device, err error) {
	d = &Device{}
	fmt.Printf("%v\n", c)
	if c.Width == 0 || c.Height == 0 {
		c.Width = 320
		c.Height = 180
		if c.Scale == 0 {
			c.Scale = 2
		}
	}
	if c.Scale == 0 {
		c.Scale = 1
	}
	w := int32(c.Width)
	h := int32(c.Height)
	s := int32(c.Scale)
	d.window, err = sdl.CreateWindow(c.Title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		w*s, h*s, sdl.WINDOW_SHOWN)
	if err == nil {
		d.render, err = sdl.CreateRenderer(d.window, -1, sdl.RENDERER_ACCELERATED)
	}
	if err == nil {
		d.render.SetLogicalSize(w, h)
		d.render.SetDrawColor(39, 58, 93, 255)
		d.render.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	}
	return
}

func (d *Device) Close() {
	if d.render != nil {
		d.render.Destroy()
	}
	if d.window != nil {
		d.window.Destroy()
	}
	sdl.Quit() //todo sdl quit only whene everything is close
	// => have a global sdl handlig 'class'
}

func (d *Device) Clear() {
	d.render.Clear()
}
