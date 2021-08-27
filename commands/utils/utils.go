package utils

import (
	"github.com/fenwikk/Tar/commands/utils/embed"
	"github.com/fenwikk/scp"
)

var Category = &scp.Category{
	Name:         "Utils",
	ID:           "utils",
	Description:  "Helpful Utilities",
	HelpEmoji:    "🛠️",
	RegisterCmds: RegisterCommands,
}

func RegisterCommands(c *scp.Category) {
	c.AddCommand(embed.Command)
}
