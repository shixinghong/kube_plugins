package main

import (
	"context"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kube_plugins/help"
	"os"
)

var clientSet = help.Client()

func main() {
	help.RunCmd(run)
}

func run(c *cobra.Command, args []string) error {
	//clientSet := help.Client()
	ns, err := c.Flags().GetString("namespace")
	if err != nil {
		return err
	}
	if ns == "" {
		ns = "default"
	}
	podList, err := clientSet.CoreV1().Pods(ns).
		List(context.Background(), metav1.ListOptions{
			LabelSelector: help.Labels,
			FieldSelector: help.Fields,
		})
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	header := []string{"名称", "命名空间", "IP", "状态"}
	if help.ShowLabels {
		header = append(header, "标签")
	}
	table.SetHeader(header)

	for _, pod := range podList.Items {
		podRow := []string{pod.Name, pod.Namespace, pod.Status.PodIP, string(pod.Status.Phase)}
		if help.ShowLabels {
			podRow = append(podRow, help.Map2String(pod.Labels))
		}
		table.Append(podRow)
	}

	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)
	// 渲染
	table.Render()
	return nil
}
