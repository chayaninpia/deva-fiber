package router

import (
	"github.com/chayaninpia/deva-fiber/src/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func InitRouter(f *fiber.App) {

	f.Get("/power/:column/:numberRecord", handler.Power)

	f.Listen(viper.GetString("server.port"))
}
