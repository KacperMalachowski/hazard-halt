package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type MaliciousPage struct {
	RegisterPositionId int64  `json:"RegisterPositionId"`
	DomainAdress       string `json:"DomainAdress"`
	InsertDate         string `json:"InsertDate"`
	DeleteDate         string `json:"DeleteDate"`
}

var pages []MaliciousPage

func main() {
	var err error
	pages, err = loadMaliciousPages("domains.json")
	if err != nil {
		panic(err)
	}

	token := os.Getenv("DISCORD_BOT_TOKEN")

	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	bot.AddHandler(messageCreate)

	bot.Identify.Intents = discordgo.IntentsGuildMessages

	err = bot.Open()
	if err != nil {
		panic(err)
	}

	fmt.Println("Bot is running...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	bot.Close()
}

func loadMaliciousPages(path string) ([]MaliciousPage, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var maliciousPages []MaliciousPage
	err = json.NewDecoder(file).Decode(&maliciousPages)
	if err != nil {
		return nil, err
	}

	return maliciousPages, nil
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Delete message if it contains a malicious page
	if containsMaliciousPage(m.Content, pages) {
		err := s.ChannelMessageDelete(m.ChannelID, m.ID)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func containsMaliciousPage(content string, mps []MaliciousPage) bool {
	for _, mp := range mps {
		if mp.DeleteDate != "" {
			continue
		}

		if strings.Contains(content, mp.DomainAdress) {
			return true
		}
	}
	return false
}
