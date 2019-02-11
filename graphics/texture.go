package graphics

import (
	"io"
	"os"
)

type Texture interface {
}

func Load(name string) (tex *Texture, err error) {
	file, err := os.Open(name)
	defer file.Close()
	if err == nil {
		tex, err = LoadFromReader(file)
	}
	return
}

func LoadFromReader(r io.Reader) (tex *Texture, err error) {
	/*img, _, err := image.Decode(r)
	if err == nil {*/
	//bounds := img.Bounds()
	/*rect.W = int32(bounds.Max.X)
	rect.H = int32(bounds.Max.Y)*/
	//	}
	/*tex, err = render.CreateTexture(uint32(sdl.PIXELFORMAT_RGBA32), sdl.TEXTUREACCESS_STATIC, rect.W, rect.H)
	if err != nil {
		panic(err)
	}
	*/
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
		draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)*/
	/*f, err := os.Create("test.png")
	png.Encode(f, rgba)
	f.Close()*/
	/*	tex.Update(&rect, rgba.Pix, rgba.Stride)
		tex.SetBlendMode(sdl.BLENDMODE_BLEND)*/

	return
}
