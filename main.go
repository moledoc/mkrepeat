package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	fN := flag.Int("n", 10, "repeat audio n times")
	flag.Parse()
	argv := flag.Args()
	if len(argv) != 1 {
		fmt.Fprintf(os.Stderr, "please provide only a single argument\n")
		os.Exit(1)
	}
	if _, err := os.Stat(argv[0]); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "provided file doesn't exist\n")
		os.Exit(1)
	}
	fp, err := filepath.Abs(argv[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get absolute path for %q: %v\n", argv[0], err)
		os.Exit(1)
	}
	var args []string
	args = append(args, "-y")
	for i := 0; i < (*fN); i++ {
		args = append(args, "-i")
		args = append(args, fp)
	}
	args = append(args, "-filter_complex")
	args = append(args, fmt.Sprintf("concat=n=%v:v=0:a=1[a]", *fN))
	args = append(args, "-map")
	args = append(args, "[a]")
	args = append(args, "-codec:a")
	args = append(args, "libmp3lame")
	args = append(args, "-b:a")
	args = append(args, "256k")
	filename := filepath.Base(argv[0])
	dirPath := strings.Replace(argv[0], filename, "", 1)
	args = append(args, fmt.Sprintf("%v/repeat - %s", dirPath, filename))
	fmt.Fprintf(os.Stdout, "composing %vx repeat of %q \n", *fN, fp)
	out, err := exec.Command("ffmpeg", args...).CombinedOutput()
	fmt.Fprintf(os.Stdout, "%s\n", out)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error when running the command: %v\n", err.Error())
		os.Exit(1)
	}
}
