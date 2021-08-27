package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/fenwikk/Tar/commands/general"
	"github.com/fenwikk/Tar/commands/utils"
	"github.com/fenwikk/scp"
)

// Bot parameters
var (
	GuildID        = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	BotToken       = flag.String("token", "", "Bot access token")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

var (
	s *discordgo.Session
)

func init() { flag.Parse() }

func main() {
	var err error
	s, err = discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	router := scp.Create(s)

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot ready")
	})
	err = s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	router.AddCategory(general.Category)
	router.AddCategory(utils.Category)

	router.RegisterAllCommands(*GuildID)

	defer s.Close()

	log.Println("Bot is up!")

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Gracefully shutdowning")
	if *RemoveCommands {
		cmds, _ := s.ApplicationCommands(s.State.User.ID, *GuildID)
		for _, c := range cmds {
			log.Println("Removing", c.Name)
			err := s.ApplicationCommandDelete(s.State.User.ID, *GuildID, c.ID)
			if err != nil {
				log.Panicln("Error removing command:", err)
			}
		}
	}
}
