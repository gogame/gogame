package graphics

import "github.com/veandco/go-sdl2/sdl"

type Device struct {
	window *sdl.Window
	render *sdl.Renderer
	//	title string
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

func NewDevice() (d *Device, err error) {
	d = &Device{}
	d.window, err = sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		640, 360, sdl.WINDOW_SHOWN)
	if err == nil {
		d.render, err = sdl.CreateRenderer(d.window, -1, sdl.RENDERER_ACCELERATED)
	}
	if err == nil {
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
