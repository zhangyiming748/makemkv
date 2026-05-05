package main

import (
	"fmt"
	"os"
	"time"

	"io"
	"log"
	"makemkv/core"

	"github.com/zhangyiming748/lumberjack"

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
	SetLog("m2ts2mkv.log")
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
			keep, _ := cmd.Flags().GetBool("keep")
			return core.M2TS2MKV(dir, flac, keep)
		},
	}

	// 添加 -d/--dir 参数，并设置为必需
	m2mCmd.Flags().StringP("dir", "d", "", "Root directory to search for M2TS files (required)")
	m2mCmd.MarkFlagRequired("dir")

	// 添加 --flac 参数，控制是否将音频转换为 FLAC 格式
	m2mCmd.Flags().BoolP("flac", "f", false, "Convert audio to FLAC format (default: false)")

	// 添加 -k/--keep 参数，控制是否保留原始 m2ts 文件
	m2mCmd.Flags().BoolP("keep", "k", false, "Keep original M2TS files after conversion (default: false)")

	// //将子命令添加到主命令
	rootCmd.AddCommand(m2mCmd)

	// 添加 flac 子命令，用于将 MKV 文件中的音频转换为 FLAC
	var flacCmd = &cobra.Command{
		Use:   "flac",
		Short: "Convert audio in MKV files to FLAC format",
		Long:  "Convert all audio streams in MKV files to FLAC format while keeping video and subtitles unchanged",
		RunE: func(cmd *cobra.Command, args []string) error {
			dir, _ := cmd.Flags().GetString("dir")
			if dir == "" {
				return fmt.Errorf("directory path cannot be empty")
			}
			return core.Mkv2Flac(dir)
		},
	}

	// 添加 -d/--dir 参数，并设置为必需
	flacCmd.Flags().StringP("dir", "d", "", "Root directory to search for MKV files (required)")
	flacCmd.MarkFlagRequired("dir")

	// 将 flac 子命令添加到主命令
	rootCmd.AddCommand(flacCmd)

	// 执行命令
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func SetLog(l string) {
	// 设置全局时区为Asia/Shanghai
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Printf("Failed to load timezone Asia/Shanghai: %v", err)
	} else {
		time.Local = location
	}
	// 创建一个用于写入文件的Logger实例
	fileLogger := &lumberjack.Logger{
		Filename:   l,
		MaxSize:    1, // MB
		MaxBackups: 1,
		MaxAge:     28, // days
	}
	err = fileLogger.Rotate()
	if err != nil {
		log.Println("Failed to rotate log file", err)
	}
	consoleLogger := log.New(os.Stdout, "CONSOLE: ", log.LstdFlags)
	log.SetOutput(io.MultiWriter(fileLogger, consoleLogger.Writer()))
	log.SetFlags(log.Ltime | log.Lshortfile)
}
