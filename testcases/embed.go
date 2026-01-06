package testcases

import (
	"embed"
)

// all:ff/0c/*.white
// 命令执行代码执行 已知漏洞利用 投递白样本 投递黑客工具 文件包含 目录遍历
// 已知漏洞利用/*/*.black 投递白样本/*/*.black 投递黑客工具/*/*.black 文件包含/*/*.black 目录遍历/*/*.black
//
// all:SQL注入/*/*.black SSRF服务端请求伪造/*/*.black Webshell上传/*/*.black XML外部实体注入/*/*.black XSS跨站脚本/*/*.black 反序列化漏洞/*/*.black 命令执行代码执行/*/*.black 已知漏洞利用/*/*.black 投递白样本/*/*.white 投递黑客工具/*/*.black 文件包含/*/*.black 目录遍历/*/*.black
//
//go:embed all:*/*/*.white all:*/*/*.black
var EmbedTestCasesFS embed.FS
