package general

import (
	"github.com/fenwikk/Tar/commands/general/info"
	"github.com/fenwikk/Tar/scp"
)

var Category = &scp.Category{
	Name:         "General",
	ID:           "general",
	Description:  "General & informational commands",
	HelpEmoji:    "ℹ️",
	RegisterCmds: RegisterCommands,
}

func RegisterCommands(c *scp.Category) {
	c.AddCommand(info.Command)
}
