package gogame

import "github.com/gogame/gogame/graphics"

type Game struct {
	config         *graphics.Config
	GraphicsDevice *graphics.Device
}

/*
func Title(g *Game) error { return g.SetTitle()}
*/
// Title sets the title of the game
// Can't be changed after startup
/*func Title(t string) func(*Game) error {
	return func(g *Game) error {
		g.title = t
		return nil
	}
}*/

func (g *Game) Dump() {
	println("Title: " + g.config.Title)
}

/*func Run(update func(*Game) error, options ...func(*Game) error) error {
	g := &Game{}
	for _, option := range options {
		option(g)
		//todo handle err and cleabup
	}
	update(g)
	//todo user ticker to recall update
	return nil
}*/

func Run(update func(*Game) error, c *graphics.Config) error {
	g := &Game{}
	g.config = c
	g.GraphicsDevice, _ = graphics.NewDevice(g.config)
	//todo handle error
	update(g)
	//todo user ticker to recall update
	return nil
}
