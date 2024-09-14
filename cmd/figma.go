package cmd

import (
	"fmt"
	"sync"

	"github.com/otfot/fasterx/internal"
	"github.com/otfot/fasterx/pkg"
	"github.com/spf13/cobra"
)

var data = []pkg.Source{
	{
		Domain: "www.figma.com",
		URL:    "https://www.figma.com/api/community_categories/all?page_size=10",
	},
	{
		Domain: "static.figma.com",
		URL:    "https://static.figma.com/app/icon/1/icon-192.png",
	},
	{
		Domain: "s3-alpha-sig.figma.com",
		URL:    "https://s3-alpha.figma.com/profile/9b3f693e-0677-4743-89ff-822b9f6b72be",
	},
}

// figmaCmd represents the figma command
var figmaCmd = &cobra.Command{
	Use:   "figma",
	Short: "获取 Figma 的优选 IP",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			result = make([]*pkg.Record, len(data))
			g      sync.WaitGroup
		)
		for i, src := range data {
			g.Add(1)
			go func() {
				defer g.Done()
				if record := internal.GetRecord(src, pkg.InDNS); record != nil {
					fmt.Printf("%s 找到优选 IP，速度 %d ms\n", src.Domain, record.Result.Duration.Milliseconds())
					result[i] = record
				} else {
					fmt.Printf("%s 域名未能成功选择到合适 IP ，请重试", src.Domain)
				}
			}()
		}
		g.Wait()

		for _, r := range result {
			if r != nil {
				fmt.Println(r.Output())
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(figmaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// figmaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// figmaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
