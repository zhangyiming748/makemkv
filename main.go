package main

import (
	"fmt"
	"os"

	"makemkv/core"

	"github.com/spf13/cobra"
)

/*
这里使用cobra实现一个命令行工具
主命令叫做makemkv
子命令为m2m 表示m2ts to mkv
功能是将m2ts文件转换为mkv文件
函数已经在core包中实现M2TS2MKV(root string)
命令中使用 -d --dir 接收root参数
参数不能为空或省略
用cobra自带的限制功能来实现
*/

func main() {
	var rootCmd = &cobra.Command{
		Use:   "makemkv",
		Short: "MakeMKV - M2TS to MKV converter",
		Long:  "A command line tool to convert M2TS files to MKV format",
	}

	var m2mCmd = &cobra.Command{
		Use:   "m2m",
		Short: "Convert M2TS files to MKV format",
		Long:  "Convert all M2TS files in the specified directory and its subdirectories to MKV format",
		RunE: func(cmd *cobra.Command, args []string) error {
			dir, _ := cmd.Flags().GetString("dir")
			if dir == "" {
				return fmt.Errorf("directory path cannot be empty")
			}
			flac, _ := cmd.Flags().GetBool("flac")
			return core.M2TS2MKV(dir, flac)
		},
	}

	// 添加 -d/--dir 参数，并设置为必需
	m2mCmd.Flags().StringP("dir", "d", "", "Root directory to search for M2TS files (required)")
	m2mCmd.MarkFlagRequired("dir")

	// 添加 --flac 参数，控制是否将音频转换为 FLAC 格式
	m2mCmd.Flags().BoolP("flac", "f", false, "Convert audio to FLAC format (default: false)")

	// 将子命令添加到主命令
	rootCmd.AddCommand(m2mCmd)

	// 执行命令
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
