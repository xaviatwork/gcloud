package gcloud

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func PrintAccessToken() (string, error) {
	out, err := gcloud("auth", "print-access-token")
	if err != nil {
		return "", err
	}
	return out, nil
}

func PrintIdentityToken() (string, error) {
	out, err := gcloud("auth", "print-identity-token")
	if err != nil {
		return "", err
	}
	return out, nil
}

func Project() (string, error) {
	out, err := gcloud("config", "get", "project")
	if err != nil {
		return "", err
	}
	return out, nil
}

func Region() (string, error) {
	out, err := gcloud("config", "get", "compute/region")
	if err != nil {
		return "", err
	}
	return out, nil
}

func Account() (string, error) {
	out, err := gcloud("config", "get", "account")
	if err != nil {
		return "", err
	}
	return out, nil
}

func gcloud(args ...string) (string, error) {
	var (
		cmdOut bytes.Buffer
		cmdErr bytes.Buffer
	)
	cmd := exec.Command("gcloud", args...)
	// Grab stdout and stderr
	cmd.Stderr = &cmdErr
	cmd.Stdout = &cmdOut
	// If there was an error, combine both errors:
	// - Why the command failed, returned from the OS
	// - Output provided by the command itself
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%s %s", err, cmdErr.String())
	}
	// Remove newline character from the command's output
	return strings.TrimSuffix(cmdOut.String(), "\n"), nil
}
