package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(
		GenerateCmd,
	)
	GenerateCmd.Flags().StringP("input", "i", "", "生成に使用するJSONファイルのパス")
	GenerateCmd.Flags().StringP("output", "o", "", "生成されたTSファイルの保存先ディレクトリのパス")
}

func Execute() {
	rootCmd.Execute()
}
