package help

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"log"
)

var cfgFlags *genericclioptions.ConfigFlags
var ShowLabels bool
var Labels string
var Fields string

func Client() *kubernetes.Clientset {
	cfgFlags = genericclioptions.NewConfigFlags(true)
	config, err := cfgFlags.ToRawKubeConfigLoader().ClientConfig()
	if err != nil {
		log.Fatal(err)
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	return clientSet
}

// MergeFlags 合并kubectl的flag
func MergeFlags(cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		cfgFlags.AddFlags(cmd.Flags())
	}
}

func RunCmd(f func(c *cobra.Command, args []string) error) {
	cmd := &cobra.Command{
		Use:          "kubectl pods [flags]",
		Short:        "list pods",
		Example:      "kubectl pods [flags]",
		SilenceUsage: true,
		RunE:         f,
	}

	cmd.Flags().BoolVar(&ShowLabels, "show-labels", false, "--show-labels")
	cmd.Flags().StringVar(&Labels, "labels", "", "kubectl pods --labels=\"k1=v1,k2=v1\"")
	cmd.Flags().StringVar(&Fields, "fields", "", "kubectl pods --fields=\"status.phase=Running\"")
	MergeFlags(cmd)
	err := cmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}

func Map2String(m map[string]string) (ret string) {
	for k, v := range m {
		ret += fmt.Sprintf("%s:%s,", k, v)
	}
	return ret
}
