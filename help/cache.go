package help

import (
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/labels"
	"os"
)

var cacheCmd = &cobra.Command{
	Use:   "cache", // 命令参数
	Short: "pods by cache",
	//Example:      "kubectl pods [flags]",
	Hidden: true,
	//SilenceUsage: true,
	RunE: func(c *cobra.Command, args []string) error {
		ns, err := c.Flags().GetString("namespace")
		if err != nil {
			return err
		}
		if ns == "" {
			ns = "default"
		}

		podList, err := fact.Core().V1().Pods().Lister().Pods(ns).List(labels.Everything())
		if err != nil {
			return err
		}

		table := tablewriter.NewWriter(os.Stdout)
		header := []string{"名称", "命名空间", "IP", "状态"}
		table.SetHeader(header)

		for _, pod := range podList {
			podRow := []string{pod.Name, pod.Namespace, pod.Status.PodIP, string(pod.Status.Phase)}
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
	},
}
