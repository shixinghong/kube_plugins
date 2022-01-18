package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"kube_plugins/help"
	"log"
)

func main() {
	clientSet := help.Client()
	cmd := &cobra.Command{
		Use:          "kubectl pods [flags]",
		Short:        "list pods",
		Example:      "kubectl pods [flags]",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			ns, err := c.Flags().GetString("namespace")
			if err != nil {
				return err
			}
			if ns == "" {
				ns = "default"
			}
			podList, err := clientSet.CoreV1().Pods(ns).List(context.Background(), metav1.ListOptions{})
			if err != nil {
				return err
			}
			for _, pod := range podList.Items {
				fmt.Println(pod.Name)
			}
			return nil
		},
	}
	help.MergeFlags(cmd)
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var cfgFlags *genericclioptions.ConfigFlags

func MergeFlags(cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		cfgFlags.AddFlags(cmd.Flags())
	}
}
