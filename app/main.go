package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	app.Handle("GET", "/", func(cx iris.Context) {
		cx.JSON(iris.Map{"message": "ping"})
	})

	app.Listen(":8081")
}
