package info

import (
	"github.com/bwmarrin/discordgo"
	"github.com/fenwikk/scp"
)

var Command = &scp.Command{
	Name:        "info",
	Description: "Gives information about the bot",
	Handler: func(ctx *scp.Ctx) {
		ctx.Respond(4, &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "ℹ️ Info",
					Description: "Just a normal, crappy bot made by `Tarsier#6520` because he didn't have anything to do",
				},
			},
		})
	},
}
