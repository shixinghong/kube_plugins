package help

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"log"
)

var cfgFlags *genericclioptions.ConfigFlags

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
