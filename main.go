package main

import (
	"log"
	"os"
	"voltguard/bot"

	"github.com/joho/godotenv"
)

func main() {

	// Load The Env Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not Load .env File. Please provide a valid .env File")
	}
	discordBotToken := os.Getenv("DISCORD_BOT_TOKEN")
	discordAppId := os.Getenv("APP_ID")
	discordGuildId := os.Getenv("GUILD_ID")

	// Start the Discord Bot

	bot.DiscordBotToken = discordBotToken
	bot.DiscordAppId = discordAppId
	bot.DiscordGuildId = discordGuildId
	bot.Run()
}
