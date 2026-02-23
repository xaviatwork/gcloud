package gcloud

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// > Requires having 'gcloud' installed and present in $PATH
// runs 'gcloud auth print-access-token'
// returns the output of the command
// or an empty string and the error (it there was an error)
func PrintAccessToken() (string, error) {
	out, err := gcloud("auth", "print-access-token")
	if err != nil {
		return "", err
	}
	return out, nil
}

// > Requires having 'gcloud' installed and present in $PATH
// runs gcloud auth print-identity-token
// returns the output of the command
// or an empty string and the error (it there was an error)
func PrintIdentityToken() (string, error) {
	out, err := gcloud("auth", "print-identity-token")
	if err != nil {
		return "", err
	}
	return out, nil
}

// > Requires having 'gcloud' installed and present in $PATH
// runs 'gcloud config get project'
// returns the output of the command
// or an empty string and the error (it there was an error)
func Project() (string, error) {
	out, err := gcloud("config", "get", "project")
	if err != nil {
		return "", err
	}
	return out, nil
}

// > Requires having 'gcloud' installed and present in $PATH
// runs 'gcloud config get compute/region'
// returns the output of the command
// or an empty string and the error (it there was an error)
func Region() (string, error) {
	out, err := gcloud("config", "get", "compute/region")
	if err != nil {
		return "", err
	}
	return out, nil
}

// > Requires having 'gcloud' installed and present in $PATH
// runs 'gcloud config get account'
// returns the output of the command
// or an empty string and the error (it there was an error)
func Account() (string, error) {
	out, err := gcloud("config", "get", "account")
	if err != nil {
		return "", err
	}
	return out, nil
}

// > Requires having 'gcloud' installed and present in $PATH
// runs the 'gcloud' executable with the provided parameters
// and returns the output of the command
// or an empty string and the error (it there was an error)
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
