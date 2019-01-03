package main

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/png"
	"os"

	"github.com/gogame/gogame"
	"github.com/gogame/gogame/graphics"
	"github.com/veandco/go-sdl2/sdl"
)

//"fmt"

/*
fct: stateless
method with value receiver: use state of struct but don't change it
method with pointer receiver: use state of struct and change it

pour chque paraemttre voir s'il est utilise à chaque appel ou pas
si non => variable

ne pas se baser sur xna
api tres simple puis étendre
*/

func main() {
	//gogame.Run(update, gogame.Title("blabla"))
	//
	//illogic que title soit dans graphics config

	//g.Load mais besoin d'instance de Game
	//besoin de ui avant load pour progress
	///m ethod start? ou load
	/*
		gogame.Start(start) ou content.Load
	*/

	gogame.Run(update, &graphics.Config{Title: "AGame", Scale: 1})
	/*Width:  640,
	Height: 360,
	Title:  "bliblo"})*/
}

func update(g *gogame.Game) error {

	g.Dump()

	toto(g.GraphicsDevice)

	return nil
}

func toto(gd *graphics.Device) {
	/*

		init game size or use default values
		spr := gg.Load('sprite.png') //ou Content.Load retourne plutot un sprite qu'une image/texture
		ou cm.Load (cm =  content manager)
		si retourne sprite, comment specifier frame rate  or frame duration etc...
		Load always returns a sprite ou music etc...; LoadTexture loads onlys texture
		have multiple load functions; load switch between then according to extension

		spr.Draw(10,10) ou gg.Draw(spr, 10, 10) draw sp'applique au srite donc certainement spr.Draw
		ou sprite contient position

		gg.Flip() ou SwapBuffer() ou trouver autre nom

	*/
	/*
		if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
			panic(err)
		}
		defer sdl.Quit()

		window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
			800, 600, sdl.WINDOW_SHOWN)
		if err != nil {
			panic(err)
		}
		defer window.Destroy()

		Render(), err := sdl.CreateRender()er(window, -1, sdl.RENDER()ER_ACCELERATED)
		if err != nil {
			panic(err)
		}*/

	/*gd, err := graphics.NewDevice(nil)
	defer gd.Close()
	if err != nil {
		panic(err)
	}*/

	/*surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	rect := sdl.Rect{0, 0, 200, 200}
	surface.FillRect(&rect, 0xffff0000)
	window.UpdateSurface()*/
	/*surface, err := sdl.CreateRGBSurface(0, 200, 200, 32, 0, 0, 0, 0)
	if err != nil {
		panic(err)
	}*/
	rect := sdl.Rect{0, 0, 200, 200}
	/*surface.FillRect(&rect, 0xffff0000)
	tex, err := Render().CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}
	surface.Free()*/
	file, err := os.Open("aak11.png")
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(file)
	file.Close()
	if err != nil {
		panic(err)
	}
	bounds := img.Bounds()
	fmt.Printf("%v\n", bounds)
	rect.W = int32(bounds.Max.X)
	rect.H = int32(bounds.Max.Y)

	tex, err := gd.Render().CreateTexture(uint32(sdl.PIXELFORMAT_RGBA32), sdl.TEXTUREACCESS_STATIC, rect.W, rect.H)
	if err != nil {
		panic(err)
	}

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
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)
	/*f, err := os.Create("test.png")
	png.Encode(f, rgba)
	f.Close()*/
	tex.Update(&rect, rgba.Pix, rgba.Stride)
	tex.SetBlendMode(sdl.BLENDMODE_BLEND)

	//gd.Render().SetDrawColor(39, 58, 93, 255)
	//gd.Render().SetDrawBlendMode(sdl.BLENDMODE_BLEND)

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
		gd.Clear()
		gd.Render().Copy(tex, &rect, &rect)
		gd.Render().Present()
	}

	tex.Destroy()
	//Render().Destroy()
	//window.Destroy()
}

/*
type Texture struct {
	foo int
}

func (t Texture) String() (string) {
	return "TEX"
}

func (t Texture) TextureFct() {
	println("TEXFCT")
}

type Sprite struct {
	Texture
	foo int
}

func (t Sprite) String() (string) {
	return "SPR"
}

func (t Texture) SpriteFct() {
	println("SPRFCT")
}

func foobar(b bool) (interface{}) {
    if  (b) {
return Texture{}
    }else {return Sprite{}}
}

func main() {
    a := foobar(true).(Texture)
    b := foobar(false).(Sprite)
    fmt.Printf("%v %T \n", a, a)
    fmt.Printf("%v %T \n", b , b)
    a.TextureFct()
    b.SpriteFct()
    b.TextureFct()
}


https://hunterloftis.github.io/2017/07/11/useful-constructors/
https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
https://gamedev.stackexchange.com/questions/151877/handling-variable-frame-rate-in-sdl2

*/
//https://github.com/oskca/gopherjs-canvas
//import "github.com/veandco/go-sdl2/sdl"
//todo allow passing argument and returning value
//argc int, argv []string
