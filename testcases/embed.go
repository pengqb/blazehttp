package testcases

import (
	"embed"
)

//go:embed all:12/*/*.white all:12/*/*.black
var EmbedTestCasesFS embed.FS
