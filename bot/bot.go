package bot

import (
	"fmt"
	"log"
	"os"
	"voltguard/hue"

	"github.com/bwmarrin/discordgo"
)

// Store API Tokens

var (
	DiscordBotToken string
	DiscordAppId    string
	DiscordGuildId  string
)

func interactionCreate(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	hueUsername := os.Getenv("HUE_USERNAME")
	hueBridgeIp := os.Getenv("HUE_BRIDGE_IP")
	bridge := hue.Connect(hueUsername, hueBridgeIp)
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
		case "lightoff":
			go func() {
				response := &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: hue.LightsOff(bridge),
					},
				}
				_ = session.InteractionRespond(interaction.Interaction, response)
			}()
		case "lighton":
			go func() {
				response := &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: hue.LightsOn(bridge),
					},
				}
				_ = session.InteractionRespond(interaction.Interaction, response)
			}()
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
		{
			Name:        "lightoff",
			Description: "Turn The lights Off",
		},
		{
			Name:        "lighton",
			Description: "Turn on The Lights",
		},
	}

	go func() {
		for i := 0; i < len(commands); i++ {
			_, err = discord.ApplicationCommandCreate(discord.State.User.ID, DiscordGuildId, commands[i])
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	fmt.Println("VoltGuard is now running. Press Ctrl+C to exit.")
	select {}
}
