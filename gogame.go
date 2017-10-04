package gogame

//import "github.com/veandco/go-sdl2/sdl"

//delta
var DeltaTime float32

//game
//type Game struct {
//	deltaTime float32
//

//hjjkh
//type GameFunc func()

//jkjk
type Game interface {
	Start()
	Update()
}

type GameFixedUpdate interface {
	FixedUpdate()
}

//run
func Run(g Game) {
	g.Start()
	g.Update()
	if c, ok := g.(GameFixedUpdate); ok {
		c.FixedUpdate()
	}

	/*if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	rect := sdl.Rect{0, 0, 200, 200}
	surface.FillRect(&rect, 0xffff0000)
	window.UpdateSurface()

	sdl.Delay(2500)*/
}