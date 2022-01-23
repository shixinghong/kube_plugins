package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"os"
	"strings"
)

func main() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(">>>"),
	)
	p.Run()
}

func executor(in string) {
	in = strings.TrimSpace(in)
	blocks := strings.Split(in, " ")
	switch blocks[0] {
	case "exit":
		fmt.Println("Exit!")
		os.Exit(0)
	}

}

var suggest = []prompt.Suggest{
	{"get", "获取pod详情"},
	{"exit", "Exit prompt"},
}

func completer(d prompt.Document) []prompt.Suggest {
	w := d.GetWordBeforeCursor()
	if w == "" {
		return []prompt.Suggest{}
	}
	return prompt.FilterHasPrefix(suggest, w, true)
}
