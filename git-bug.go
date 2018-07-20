//go:generate go run pack_webui.go
//go:generate go run doc/gen_markdown.go

package main

import "github.com/MichaelMure/git-bug/commands"

func main() {
	commands.Execute()
}
