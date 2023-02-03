package main

import (
	"github.com/chayaninpia/deva-fiber/src/router"
	"github.com/chayaninpia/deva-fiber/src/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {

	utils.InitConfig()

	f := fiber.New()

	router.InitRouter(f)

}
