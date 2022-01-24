package help

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/labels"
	"log"
	"os"
	"regexp"
	"strings"
)

func executorCmd(cmd *cobra.Command) func(in string) {
	return func(in string) {
		in = strings.TrimSpace(in)
		var args []string
		blocks := strings.Split(in, " ")
		if len(blocks) > 1 {
			args = blocks[1:]
		}
		switch blocks[0] {
		case "exit":
			fmt.Println("Exit!")
			os.Exit(0)
		case "list": //列出pods
			if err := cacheCmd.RunE(cmd, args); err != nil {
				log.Fatal(err)
			}
		case "get":
			getPodDetails(args)
		}
	}
}

func getPodList() []prompt.Suggest {
	sug := make([]prompt.Suggest, 0)
	pods, err := fact.Core().V1().Pods().Lister().Pods("default").List(labels.Everything())
	if err != nil {
		return sug
	}
	for _, pod := range pods {
		sug = append(sug, prompt.Suggest{
			Text:        pod.Name,
			Description: "节点" + pod.Spec.NodeName + " 状态:" + string(pod.Status.Phase) + " IP:" + pod.Status.PodIP,
		})
	}
	return sug
}

func completer(d prompt.Document) []prompt.Suggest {
	//w := d.GetWordBeforeCursor()
	s := []prompt.Suggest{
		{"get", "获取pod详情"},
		{"exit", "Exit prompt"},
	}
	cmd, opt := parseCmd(d.TextBeforeCursor())
	if cmd == "get" {
		return prompt.FilterHasPrefix(getPodList(), opt, true)
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func parseCmd(w string) (string, string) {
	w = regexp.MustCompile("\\s+").ReplaceAllString(w, " ")
	l := strings.Split(w, " ")
	if len(l) >= 2 {
		return l[0], strings.Join(l[1:], " ")
	}
	return w, ""
}

var promptCmd = &cobra.Command{
	Use:          "prompt", // 命令参数
	Short:        "prompt pods",
	Example:      "kubectl pods prompt [flags]",
	SilenceUsage: true,
	RunE: func(c *cobra.Command, args []string) error {
		InitCache() // 初始化缓存
		p := prompt.New(
			executorCmd(c),
			completer,
			prompt.OptionPrefix(">>>"),
		)
		p.Run()
		return nil
	},
}
