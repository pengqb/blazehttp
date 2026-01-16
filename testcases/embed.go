package testcases

import (
	"embed"
)

// SQL注入291/351 SSRF服务端请求伪造17/32 Webshell上传26/57 XML外部实体注入20/36 XSS跨站脚本327/365
// 命令执行代码执行64/97 已知漏洞利用615/984 投递白样本 文件包含44/69 目录遍历53/66
// 删除post 反序列化漏洞/*/*.black 投递黑客工具/*/*.black Webshell上传/*/*.black
//
// all:SQL注入/*/*.black SSRF服务端请求伪造/*/*.black Webshell上传/*/*.black XML外部实体注入/*/*.black XSS跨站脚本/*/*.black 反序列化漏洞/*/*.black 命令执行代码执行/*/*.black 已知漏洞利用/*/*.black 投递白样本/*/*.white 文件包含/*/*.black 目录遍历/*/*.black
//
// all:*/*/*.white all:*/*/*.black
//
//go:embed all:已知漏洞利用/Acrolinx_Server_for_Windows_路径遍历漏洞(CVE-2018-7719)/*.black
var EmbedTestCasesFS embed.FS
