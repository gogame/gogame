package gogame

import "fmt"

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
/*type Game interface {
	//Start()
	Update()
}

type GameFixedUpdate interface {
	FixedUpdate()
}*/

//
type Starter interface {
	Start()
}

//
type Updater interface {
	Update()
}

//u
type UpdaterFunc func()

//a
func (f UpdaterFunc) Update() {
	if f != nil {
		f()
	}
}

//a
type Game struct {
	Start       func()
	Update      Updater
	FixedUpdate func()
}

//a
type StarterFunc func()

//a
func (f StarterFunc) Start() {
	f()
}

//a

//todo: renomer methods en updater
/*func Run(u GameUpdateFunc) {
	g := &Game{Methods: u}
	g.Run()
}*/

// Run is a simple method when only Start and Update are needed
func Run(start, update func()) {
	g := &Game{Start: start, Update: UpdaterFunc(update)}
	g.Run()
}

/*
//
func Run(u Updater) {
	g := &Game{Updater: u}
	g.Run()
}

//
func RunFunc(u func()) {
	Run(UpdaterFunc(u))
}
*/
//

//run
func (g *Game) Run() {

	//g.Start()
	if g.Start != nil {
		g.Start()
	}
	fmt.Printf("%v %T\n", g.Update, g.Update)
	if g.Update != nil {
		fmt.Println("not nil")
		g.Update()
	} else {
		fmt.Println("nil")
	}

	/*if c, ok := g.(GameFixedUpdate); ok {
		c.FixedUpdate()
	}

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

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	rect := sdl.Rect{0, 0, 200, 200}
	surface.FillRect(&rect, 0xffff0000)
	window.UpdateSurface()

	sdl.Delay(2500)*/
}
