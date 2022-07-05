package main

import (
	"tgbot/configs"
	"tgbot/controllers"
)

func main() {
	bot, updates := configs.SetupBot()

	controllers.HandleUpdates(bot, updates)
}
