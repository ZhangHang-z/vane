package down

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func GitRepDownloader(src string) bytes.Buffer {
	cmd := exec.Command("git", "clone", src)
	cmd.Stdin = strings.NewReader(src)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal("Exectued Command Failed:", err)
	}
	return out
}
