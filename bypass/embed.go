package bypass

import (
	"embed"
)

// all:BigInt绕过/*.black
//
//go:embed all:*/*.black
var EmbedBypassFS embed.FS
