package git

import (
	"log"
	"os/exec"
	"strings"
)

func GetDiffFiles() (map[string]string, error) {

	var diffFiles = make(map[string]string)

	// get staged files
	cmdStaged := exec.Command("git", "diff", "--name-only", "--cached")

	// get output of the command
	out, err := cmdStaged.Output()
	if err != nil {
		return nil, err
	}
	outs := strings.Split(
		strings.TrimSpace(string(out)),
		"\n",
	)

	for _, path := range outs {
		cmdDiff := exec.Command("git", "diff", "--staged", "--unified=0", path)
		outDiff, err := cmdDiff.Output()
		if err != nil {
			log.Printf("Error: %s", err.Error())
			continue
		}

		diffFiles[path] = string(outDiff)

	}

	return diffFiles, nil
}

func Commit(msg string) error {
	cmd := exec.Command("git", "commit", "-m", msg)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
