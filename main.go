package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"

	"go-iris-sample/setups"
)

func main() {
	app := iris.New()

	app.Use(iris.Compression)
	app.Configure(iris.WithoutBodyConsumptionOnUnmarshal)

	// ログ記録（これも備えつきのミドルウェア）
	ac := accesslog.File("./access.log")
	defer ac.Close()
	app.UseRouter(ac.Handler)
	app.UseRouter(recover.New())

	//"/user"から始まるURLを受け取った際の処理をグループにする
	users := app.Party("/users")
	mvc.Configure(users, setups.ConfigureUsersControllers)

	app.Listen(":8080")
}
