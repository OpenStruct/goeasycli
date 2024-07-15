package main

import (
	"embed"
	"goeasycli/cmd"
	"goeasycli/utils"
)

//go:embed templates
var templates embed.FS

func main() {
	utils.SetTemplatesFS(templates)
	cmd.Execute()
}
