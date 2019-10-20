package router

import (
	"fmt"
	"go-anything/src/Railway12306"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func newApp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	return app
}
func register(app *iris.Application) {
	app.PartyFunc("/12306", Railway12306.RegisterRailWay12306)

}
func Run(port string) {
	app := newApp()
	register(app)
	app.Run(iris.Addr(fmt.Sprintf(":%s", port)))
}