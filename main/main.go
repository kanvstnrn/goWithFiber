package main

import (
	"github.com/gofiber/fiber"
)

func main(){
	app := fiber.New()
	app.Get("/",helloWorld)
	_ = app.Listen(":3000")


}

func helloWorld(ctx *fiber.Ctx) error {
	_ = ctx.Send([]byte("Hello"))
	return nil
}