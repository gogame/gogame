package gogame

import (
	"github.com/gogame/gogame/video"
	"github.com/veandco/go-sdl2/sdl"
)

//delta
var DeltaTime float32

type Starter interface {
	Start()
}

type Updater interface {
	Update()
}

type FixedUpdater interface {
	Update()
}

func RunGame(u Updater) {
	if s, ok := u.(Starter); ok {
		s.Start()
	}
	u.Update()
}
/*
gogame.Game struct {
	init func()
	update() etc....
}

Mygame{
	Game
}
///utilisation

import (
	"fmt"
)


type Game struct {
	Update func()
}

func NewGame() *Game {
	g := new(Game)
	g.Update = g.DefaultUpdate
	return g
}

func (g Game) DefaultUpdate() { // or BaseUpdate or func GameUpdate()
	fmt.Println("Hello, Default")
}

func (g Game) Run() {
	g.Update()
}


func main() {
	g := NewGame()
	g.Update = func() {
		fmt.Println("Hello, Specific")
		g.DefaultUpdate()
	}
	g.Run()
	
}
-----

gg.Update(func() void {

})

initialize fct with code: prototype usge before anything

*/
//fenetre = pygame.display.set_mode((640,480), RESIZABLE)

// love.window.setMode( width, height, flags ) //fulscrene, resisable, etc..

//glview = GLViewImpl::createWithRect("SimpleGame", Rect(0,0, 480, 320), 1.0);
/*Size designSize = Size(480,320);
Size resourceSize = Size(960,640);

director->setContentScaleFactor(resourceSize.height / designSize.height);
director->getOpenGLView()->setDesignResolutionSize(
  designSize.width, designSize.height, ResolutionPolicy::FIXED_HEIGHT);

   graphics = new GraphicsDeviceManager(this);
        graphics.PreferredBackBufferHeight = 600;
		graphics.PreferredBackBufferWidth = 800;

		 graphics.IsFullScreen = userRequestedFullScreen;
            graphics.PreferredBackBufferHeight = userRequestedHeight;
            graphics.PreferredBackBufferWidth = userRequestedWidth;
            graphics.ApplyChanges();


*/
//todo make common functions for both run functions
func Run(start, update func()) {
	vid := video.ActiveDevice
	vid.Start()
	if start != nil {
		start()
	}
	//todo timed update loop
	update()
	vid.Update()
	for i := 0; i < 2500; i++ {
		sdl.PumpEvents()
		sdl.Delay(1)
	}
	vid.Exit()
}

/*func Run(u Updater) {
if s, ok := u.(Starter); ok {
	s.Start()
}
u.Update()*/
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

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	rect := sdl.Rect{0, 0, 200, 200}
	surface.FillRect(&rect, 0xffff0000)
	window.UpdateSurface()

	sdl.Delay(2500)*/
//}
