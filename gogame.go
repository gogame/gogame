package gogame

type Game interface {
	Update()
	//Graphics graphics.Device
	//Width uint
	/*, Height uint
	Title         string*/
}

/*
func (x *Game) Init() {
	println("GOGAME init")
}*/

type GameTime struct {
	t float32
}

type Initializable interface {
	Init()
}

type Updatable interface {
	Update()
}

type Drawable interface {
	Draw()
}

func Run(game Game) {

	init, ok := game.(Initializable)
	println("init: ", init, ok)
	if init != nil {
		init.Init()
	}

	upd, ok := game.(Updatable)
	println("update:", upd, ok)
	if upd != nil {
		upd.Update()
	}
}
