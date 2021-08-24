package embed

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/fenwikk/Tar/scp"
)

func CreateEmbed(ctx *scp.Ctx) {
	ctx.Respond(5, &discordgo.InteractionResponseData{})
	var m *discordgo.MessageCreate
	embed := []*discordgo.MessageEmbed{
		{
			Title: "🛠️ Create Embed",
		},
	}

	ctx.Respond(4, &discordgo.InteractionResponseData{
		Embeds: embed,
	})

	// Channel
	embed[0].Description = "What channel should I send the embed to?"
	ctx.EditResponse(&discordgo.WebhookEdit{Embeds: embed})

	m = ctx.WaitForResponse()

	for ctx.StrToID(m.Content, scp.ChannelMentionID) == "" {
		m = ctx.WaitForResponse()
	}

	channel, _ := ctx.Session.Channel(ctx.StrToID(m.Content, scp.ChannelMentionID))

	embed[0].Fields = append(embed[0].Fields, &discordgo.MessageEmbedField{
		Name:  "Channel",
		Value: channel.Mention(),
	})

	// Title
	log.Println("Title setup")
	embed[0].Description = "Roger! What should the title be?"
	ctx.EditResponse(&discordgo.WebhookEdit{Embeds: embed})

	log.Println("response")
	m = ctx.WaitForResponse()

	log.Println(embed[0].Fields)
	embed[0].Fields = append(embed[0].Fields, &discordgo.MessageEmbedField{
		Name:  "Title",
		Value: m.Content,
	})
	log.Println(embed[0].Fields)
	ctx.EditResponse(&discordgo.WebhookEdit{Embeds: embed})
}
