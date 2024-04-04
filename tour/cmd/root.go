package cmd

import (
	"github.com/spf13/cobra"
)

// 根命令
var rootCmd = &cobra.Command{}

// Execute 将所有子命令添加到root命令并适当设置标志。
// 这由 main.main() 调用。它只需要对 rootCmd 调用一次。
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// 接受参数
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
	rootCmd.AddCommand(sqlCmd)
}
