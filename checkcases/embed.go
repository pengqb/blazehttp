package checkcases

import (
	"embed"
)

//go:embed all:*/*/*.white all:*/*/*.black
var EmbedCheckCasesFS embed.FS
