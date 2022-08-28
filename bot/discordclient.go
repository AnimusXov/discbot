package bot

import (
	"discordbot/cache"
	"discordbot/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

var BotID string
var goBot *discordgo.Session

func Run() {
	// create bot session
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal(err)
		return
	}
	// make the bot a user
	user, err := goBot.User("@me")
	if err != nil {
		log.Fatal(err)
		return
	}
	BotID = user.ID
	goBot.AddHandler(messageHandler)
	err = goBot.Open()
	if err != nil {
		return
	}
}
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == BotID {
		return
	}
	var parameters = strings.Fields(m.Content)
	// Post a link to latest user's tweet
	if strings.HasPrefix(m.Content, "!latest") && len(parameters) == 2 {
		var value, found = cache.Read(parameters[1])
		if found {
			_, _ = s.ChannelMessageSend(m.ChannelID, "https://twitter.com/"+parameters[1]+"/status/"+fmt.Sprint(value))
		} else {
			tweet, _ := GetTimeline(parameters[1])
			if tweet.ID == 0 {
				_, _ = s.ChannelMessageSend(m.ChannelID, "No user with such parameters found")
				return
			}
			_, _ = s.ChannelMessageSend(m.ChannelID, "https://twitter.com/"+parameters[1]+"/status/"+tweet.IDStr)
			cache.Put(parameters[1], tweet.IDStr)
		}
	}

	if strings.HasPrefix(m.Content, "!latest forks") && len(parameters) == 4 {
		var parameters = strings.Fields(m.Content)
		var owner = parameters[2]
		var repo = parameters[3]
		client, ctx := config.CreateGithubClient()
		var repos = listAllForks(owner, repo, client, ctx)
		var embedMessages []discordgo.MessageEmbed

		for i := 0; i < len(repos); i++ {
			comp, isAhead := compareBranches(client, ctx, repos[i], owner)
			if isAhead {
				embedMessages = append(embedMessages,
					discordgo.MessageEmbed{
						URL:         repos[i].GetHTMLURL(),
						Type:        "rich",
						Title:       repos[i].GetHTMLURL(),
						Description: comp,
						Timestamp:   "",
						Color:       1752220,
						Footer:      nil,
						Thumbnail:   nil,
						Video:       nil,
						Provider:    nil,
						Author: &discordgo.MessageEmbedAuthor{
							URL:  repos[i].GetOwner().GetHTMLURL(),
							Name: repos[i].GetOwner().GetName(),
						},
						Fields: nil,
					})

			}

		}
		for _, msg := range embedMessages {
			_, err := s.ChannelMessageSendEmbed(m.ChannelID, &msg)
			if err != nil {
				return
			}
		}

	}
}
