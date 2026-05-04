package core

import (
	"fmt"
	"log"

	"os/exec"
	"path/filepath"
	"strings"
)

func m2ts2mkv(m2ts string) error {
	fmt.Println("m2ts to mkv")
	/*
			ffmpeg -i /vol2/1000/disk3/原盘/unzip/mount/BDMV/STREAM/00005.m2ts \
		-c:v copy -c:a flac -c:s copy \
		-map 0 -copy_unknown \
		-fflags +genpts -avoid_negative_ts make_zero \
		-f matroska /vol2/1000/disk3/原盘/unzip/让子弹飞.mkv
	*/
	var (
		args []string
		mkv  string
		cmd  *exec.Cmd
	)
	mkv = strings.Replace(m2ts, filepath.Ext(m2ts), ".mkv", 1)
	args = append(args, "-i", m2ts)
	args = append(args, "-c:v", "copy")
	args = append(args, "-c:a", "flac")
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
