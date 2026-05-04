package core

import (
	"fmt"
	"log"

	"os/exec"
	"path/filepath"
	"strings"

	"github.com/zhangyiming748/finder"
)

func M2TS2MKV(root string, flac bool) error {
	folders := finder.FindAllFolders(root)
	for i, folder := range folders {
		log.Printf("正在处理%d/%d\n文件夹 %s\n", i+1, len(folders), folder)
		files := finder.FindAllFiles(folder)
		for j, file := range files {
			log.Printf("正在处理%d/%d\n文件 %s\n", j+1, len(files), file)
			ext := strings.ToLower(filepath.Ext(file))
			if ext == ".m2ts" {
				if err := m2ts2mkv(file, flac); err != nil {
					log.Printf("处理文件 %s 失败：%s\n", file, err)
					continue
				}
			}
		}
	}
	return nil
}

func m2ts2mkv(m2ts string, flac bool) error {
	fmt.Println("m2ts to mkv")
	var (
		args []string
		mkv  string
		cmd  *exec.Cmd
	)
	mkv = strings.Replace(m2ts, filepath.Ext(m2ts), ".mkv", 1)
	args = append(args, "-i", m2ts)
	args = append(args, "-c:v", "copy")
	if flac {
		args = append(args, "-c:a", "flac")
	} else {
		args = append(args, "-c:a", "copy")
	}
	args = append(args, "-c:s", "copy")
	args = append(args, "-map", "0")
	args = append(args, "-copy_unknown")
	args = append(args, "-fflags", "+genpts")
	args = append(args, "-avoid_negative_ts", "make_zero")
	args = append(args, "-f", "matroska")
	args = append(args, mkv)
	cmd = exec.Command("ffmpeg", args...)
	log.Printf("准备运行的命令是%s\n", cmd.String())
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	} else {
		log.Printf("命令输出是%s\n", string(out))
		return nil
	}
}
