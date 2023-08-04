package bot

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

// Store API Tokens

var (
	DiscordBotToken string
	DiscordAppId    string
	DiscordGuildId  string
)

func interactionCreate(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	fmt.Println("I've got an Interaction")
	// Ignore all interactions from Bots
	if interaction.Type == discordgo.InteractionApplicationCommand {
		command := interaction.ApplicationCommandData()

		// Handle the specific command
		switch command.Name {
		case "hello":
			response := &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Hello you :)",
				},
			}
			_ = session.InteractionRespond(interaction.Interaction, response)
		}

	}
}

func Run() {
	// Create new Discord Session
	discord, err := discordgo.New("Bot " + DiscordBotToken)
	if err != nil {
		log.Fatal(err)
	}

	// Add event Handle for interactionCreate
	discord.AddHandler(interactionCreate)

	// Open Session
	discord.Open()

	defer discord.Close()

	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "hello",
			Description: "Basic",
		},
	}

	_, err = discord.ApplicationCommandCreate(discord.State.User.ID, DiscordGuildId, commands[0])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("VoltGuard is now running. Press Ctrl+C to exit.")
	select {}
}
