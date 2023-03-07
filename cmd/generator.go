package cmd

import (
	"microcms-js-client-generator-cli/generator"
	"path/filepath"

	"github.com/spf13/cobra"
)

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "クライアントファイルを生成します",
	RunE: func(cmd *cobra.Command, _ []string) error {
		ip, _ := cmd.Flags().GetString("input")
		op, _ := cmd.Flags().GetString("output")
		aip, _ := filepath.Abs(ip)
		aop, _ := filepath.Abs(op)
		generator.GenerateBaseClientFile(aip, aop)
		generator.GenerateClientFiles(aip, aop)

		return nil
	},
}
