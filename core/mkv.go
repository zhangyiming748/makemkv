package core

import (
	"fmt"
	"log"
	"os"

	"os/exec"
	"path/filepath"
	"strings"

	"github.com/zhangyiming748/finder"
)

func M2TS2MKV(root string, flac bool, keep bool) error {
	folders := finder.FindAllFolders(root)
	for i, folder := range folders {
		log.Printf("Processing %d/%d\nFolder: %s\n", i+1, len(folders), folder)
		files := finder.FindAllFiles(folder)
		for j, file := range files {
			log.Printf("Processing %d/%d\nFile: %s\n", j+1, len(files), file)
			ext := strings.ToLower(filepath.Ext(file))
			if ext == ".m2ts" || ext == ".ts" {
				if err := m2ts2mkv(file, flac, keep); err != nil {
					log.Printf("Failed to process file %s: %s\n", file, err)
					continue
				}
			}
		}
	}
	return nil
}

func m2ts2mkv(m2ts string, flac bool, keep bool) error {
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
	log.Printf("Command: %s\n", cmd.String())
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Output: %s\n", string(out))
		return err
	} else {

		if keep {
			return nil
		} else {
			return os.Remove(m2ts)
		}
	}
}

func Mkv2Flac(root string) error {
	folders := finder.FindAllFolders(root)
	for i, folder := range folders {
		log.Printf("正在处理%d/%d\n文件夹 %s\n", i+1, len(folders), folder)
		files := finder.FindAllFiles(folder)
		for j, file := range files {
			log.Printf("正在处理%d/%d\n文件 %s\n", j+1, len(files), file)
			ext := strings.ToLower(filepath.Ext(file))
			if ext == ".mkv" {
				if err := mkv2flac(file); err != nil {
					log.Printf("处理文件 %s 失败：%s\n", file, err)
					continue
				}
			}
		}
	}
	return nil
}

func mkv2flac(mkv string) error {
	fmt.Println("mkv to flac")
	var (
		args     []string
		tmp_name string
		cmd      *exec.Cmd
	)
	tmp_name = strings.Replace(mkv, filepath.Ext(mkv), "_tmp.mkv", 1)
	args = append(args, "-i", mkv)
	args = append(args, "-c:v", "copy")
	args = append(args, "-c:a", "flac")
	args = append(args, "-c:s", "copy")
	args = append(args, tmp_name)
	cmd = exec.Command("ffmpeg", args...)
	log.Printf("Command: %s\n", cmd.String())
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Command output: %s\n", string(out))
		return err
	} else {
		os.Remove(mkv)
		os.Rename(tmp_name, mkv)
		log.Printf("已删除原始文件 %s 并重命名 %s\n", mkv, tmp_name)
		return nil
	}
}
