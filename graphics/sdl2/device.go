package sdl2

import (
	"github.com/veandco/go-sdl2/sdl"
)

//todo: remplacer par device car ne devrait pas être exporté
//pourrais même etre placé dans un dossier internal
//ou alors sdl2 au même niveau que graphics
//open close device?
type Device struct {
	window *sdl.Window
	render *sdl.Renderer
}

func NewDevice() (d Device, err error) {
	return Device{}, nil
}

func (d Device) R() *sdl.Renderer {
	return d.render
}

func (d Device) Render(tex *sdl.Texture, rect sdl.Rect) {
	d.render.Clear()
	d.render.Copy(tex, &rect, &rect)
	d.render.Present()
}

/*func (p Device) CreateTexture(img Image)
tex, err = render.CreateTexture(uint32(sdl.PIXELFORMAT_RGBA32), sdl.TEXTUREACCESS_STATIC, rect.W, rect.H)
	if err != nil {
		panic(err)
	}*/

/*const size = 200 * 200 * 4
var pixels [size]byte
for i := 0; i < size; i++ {
	pixels[i] = 0xff
	i++
	pixels[i] = 0xff
	i++
	pixels[i] = 0x00
	i++
	pixels[i] = 0x00
}
tex.Update(&rect, pixels[:], 200*4)*/
/*	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)
*/ /*f, err := os.Create("test.png")
png.Encode(f, rgba)
f.Close()*/
/*	tex.Update(&rect, rgba.Pix, rgba.Stride)
	tex.SetBlendMode(sdl.BLENDMODE_BLEND)
}*/

func (d *Device) Start() {
	var err error
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	d.window, err = sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 320, 180, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	//defer window.Destroy()
	//if err == nil {
	d.render, err = sdl.CreateRenderer(d.window, -1, sdl.RENDERER_ACCELERATED)
	//}
	if err == nil {
		//render.SetLogicalSize(w, h)
		d.render.SetDrawColor(39, 58, 93, 255)
		d.render.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	}
}

func (d Device) Stop() {
	d.render.Destroy()
	d.window.Destroy()
	sdl.Quit()
}
