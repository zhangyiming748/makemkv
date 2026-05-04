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
			if ext == ".m2ts" {
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
		return err
	} else {
		log.Printf("Output: %s\n", string(out))
		if keep {
			return nil
		} else {
			return os.Remove(m2ts)
		}
	}
}
