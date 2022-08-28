package main

import (
	"discordbot/bot"
	"discordbot/config"
	"log"
)

func main() {
	config.CreateTwitterClient()
	err := config.ReadConfig("discordconfig")
	if err != nil {
		log.Fatal(err)
		return
	}
	bot.Run()
	<-make(chan struct{})
	return
}
