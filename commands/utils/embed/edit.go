package embed

import (
	"github.com/bwmarrin/discordgo"
	"github.com/fenwikk/Tar/scp"
)

func Edit(ctx *scp.Ctx) {
	ctx.Respond(4, &discordgo.InteractionResponseData{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       "Create",
				Description: "Just a normal, crappy bot made by `Tarsier#6520` because he didn't have anything to do",
			},
		},
	})

}
